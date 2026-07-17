import { useMemo, useReducer, useRef, type KeyboardEvent } from 'react';
import { calculateExpression } from '../services/calculatorApi';
import { validateExpression } from '../utils/validation';

type State = {
  expression: string;
  result: string | null;
  error: string | null;
  loading: boolean;
  history: string[];
};

type Action =
  | { type: 'APPEND'; value: string }
  | { type: 'DELETE' }
  | { type: 'CLEAR' }
  | { type: 'SET_LOADING'; value: boolean }
  | { type: 'SET_RESULT'; value: string }
  | { type: 'SET_ERROR'; value: string | null }
  | { type: 'SET_EXPRESSION'; value: string }
  | { type: 'ADD_HISTORY'; value: string };

const initialState: State = {
  expression: '',
  result: null,
  error: null,
  loading: false,
  history: [],
};

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case 'APPEND':
      return {
        ...state,
        expression: `${state.expression}${action.value}`,
        error: null,
      };
    case 'DELETE':
      return {
        ...state,
        expression: state.expression.slice(0, -1),
        error: null,
      };
    case 'CLEAR':
      return { ...state, expression: '', error: null, result: null };
    case 'SET_LOADING':
      return { ...state, loading: action.value };
    case 'SET_RESULT':
      return { ...state, result: action.value, loading: false };
    case 'SET_ERROR':
      return { ...state, error: action.value, loading: false };
    case 'SET_EXPRESSION':
      return { ...state, expression: action.value, error: null };
    case 'ADD_HISTORY':
      return {
        ...state,
        history: [...state.history.slice(-9), action.value],
      };
    default:
      return state;
  }
}

function normalizeExpressionForBackend(expression: string) {
  let normalized = expression.replace(/\^/g, '**');
  const compactExpression = normalized.replace(/\s+/g, '');

  const percentOfPattern = /(\d+(?:\.\d+)?)%(\d+(?:\.\d+)?)/g;
  normalized = compactExpression.replace(percentOfPattern, '$2*($1/100)');
  normalized = normalized.replace(/(\d+(?:\.\d+)?)%/g, '($1/100)');

  return normalized;
}

export function useCalculator() {
  const [state, dispatch] = useReducer(reducer, initialState);
  const historyPreview = useMemo(() => state.history.slice(-5), [state.history]);
  const cacheRef = useRef(new Map<string, string>());

  const evaluate = async () => {
    const trimmedExpression = state.expression.trim();
    const validationError = validateExpression(trimmedExpression);

    if (validationError) {
      dispatch({ type: 'SET_ERROR', value: validationError });
      return;
    }

    const normalizedExpression = trimmedExpression.replace(/\s+/g, '');
    const cachedResult = cacheRef.current.get(normalizedExpression);
    if (cachedResult) {
      dispatch({ type: 'SET_RESULT', value: cachedResult });
      dispatch({ type: 'SET_ERROR', value: null });
      dispatch({ type: 'ADD_HISTORY', value: `${trimmedExpression} = ${cachedResult}` });
      return;
    }

    dispatch({ type: 'SET_LOADING', value: true });
    try {
      const backendExpression = normalizeExpressionForBackend(trimmedExpression);
      const response = await calculateExpression(backendExpression);
      const resultText = String(response.result);
      cacheRef.current.set(normalizedExpression, resultText);
      dispatch({ type: 'SET_RESULT', value: resultText });
      dispatch({ type: 'SET_ERROR', value: null });
      dispatch({ type: 'ADD_HISTORY', value: `${trimmedExpression} = ${resultText}` });
    } catch (error) {
      dispatch({ type: 'SET_ERROR', value: error instanceof Error ? error.message : 'Unexpected error' });
    }
  };

  const handleInput = (value: string) => {
    if (value === 'C') {
      dispatch({ type: 'CLEAR' });
      return;
    }
    if (value === 'DEL') {
      dispatch({ type: 'DELETE' });
      return;
    }
    if (value === '=') {
      void evaluate();
      return;
    }
    if (value === 'sqrt(') {
      dispatch({ type: 'APPEND', value: 'sqrt(' });
      return;
    }
    dispatch({ type: 'APPEND', value });
  };

  const handleExpressionChange = (value: string) => {
    const sanitizedValue = value.replace(/[^0-9+\-*/%^().\sA-Za-z]/g, '');
    dispatch({ type: 'SET_EXPRESSION', value: sanitizedValue });
    const validationError = validateExpression(sanitizedValue);
    if (validationError) {
      dispatch({ type: 'SET_ERROR', value: validationError });
    } else {
      dispatch({ type: 'SET_ERROR', value: null });
    }
  };

  const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      event.preventDefault();
      void evaluate();
    }
    if (event.key === 'Escape') {
      dispatch({ type: 'CLEAR' });
    }
  };

  return {
    state,
    historyPreview,
    handleInput,
    handleExpressionChange,
    handleKeyDown,
  };
}
