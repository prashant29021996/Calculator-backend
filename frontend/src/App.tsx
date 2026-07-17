import { useMemo, useReducer } from 'react';
import { calculateExpression } from './services/calculatorApi';
import './styles/app.css';

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
  | { type: 'SET_EXPRESSION'; value: string };

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
        expression: state.expression + action.value,
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
    default:
      return state;
  }
}

const buttons = ['7', '8', '9', '/', '4', '5', '6', '*', '1', '2', '3', '-', '0', '.', '=', '+', '(', ')', 'sqrt(', 'C'];

function App() {
  const [state, dispatch] = useReducer(reducer, initialState);

  const historyPreview = useMemo(() => state.history.slice(-5), [state.history]);

  const evaluate = async () => {
    if (!state.expression.trim()) {
      dispatch({ type: 'SET_ERROR', value: 'Expression is required' });
      return;
    }

    dispatch({ type: 'SET_LOADING', value: true });
    try {
      const response = await calculateExpression(state.expression);
      dispatch({ type: 'SET_RESULT', value: String(response.result) });
      dispatch({ type: 'SET_ERROR', value: null });
      dispatch({ type: 'SET_EXPRESSION', value: state.expression });
      dispatch({ type: 'SET_LOADING', value: false });
      if (state.expression) {
        state.history.push(`${state.expression} = ${response.result}`);
      }
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
    dispatch({ type: 'APPEND', value });
  };

  return (
    <div className="app-shell">
      <div className="calculator-card">
        <h1>Calculator</h1>
        <label className="sr-only" htmlFor="expression">Expression</label>
        <input
          id="expression"
          value={state.expression}
          onChange={(event) => dispatch({ type: 'SET_EXPRESSION', value: event.target.value })}
          placeholder="Enter an expression"
          className="display"
          aria-label="Expression"
        />
        <div className="result-panel" aria-live="polite">
          {state.loading ? <span>Calculating…</span> : state.result !== null ? <strong>{state.result}</strong> : <span>Result will appear here</span>}
        </div>
        {state.error ? <p className="error-text">{state.error}</p> : null}
        <div className="keypad" role="grid">
          {buttons.map((button) => (
            <button
              key={button}
              className={`key ${button === '=' ? 'operator' : ''}`}
              onClick={() => handleInput(button)}
              type="button"
            >
              {button}
            </button>
          ))}
        </div>
        <div className="history-panel">
          <h2>Recent calculations</h2>
          {historyPreview.length ? (
            <ul>
              {historyPreview.map((entry) => (
                <li key={entry}>{entry}</li>
              ))}
            </ul>
          ) : (
            <p>No history yet.</p>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
