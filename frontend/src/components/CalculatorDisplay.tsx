import type { KeyboardEvent } from 'react';

type CalculatorDisplayProps = {
  value: string;
  onChange: (value: string) => void;
  onKeyDown: (event: KeyboardEvent<HTMLInputElement>) => void;
  result: string | null;
  loading: boolean;
  error: string | null;
};

function CalculatorDisplay({ value, onChange, onKeyDown, result, loading, error }: CalculatorDisplayProps) {
  return (
    <>
      <label className="sr-only" htmlFor="expression">
        Expression
      </label>
      <input
        id="expression"
        value={value}
        onChange={(event) => onChange(event.target.value)}
        onKeyDown={onKeyDown}
        placeholder="Enter an expression"
        className="display"
        aria-label="Expression"
        inputMode="text"
        autoComplete="off"
        spellCheck={false}
      />
      <div className="result-panel" aria-live="polite" role="status">
        {loading ? <span>Calculating…</span> : result !== null ? <strong>{result}</strong> : <span>Result will appear here</span>}
      </div>
      {error ? <p className="error-text">{error}</p> : null}
    </>
  );
}

export default CalculatorDisplay;
