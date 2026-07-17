import CalculatorDisplay from './components/CalculatorDisplay';
import HistoryPanel from './components/HistoryPanel';
import Keypad from './components/Keypad';
import { CALCULATOR_BUTTONS } from './constants/calculator';
import { useCalculator } from './hooks/useCalculator';
import './styles/app.css';

function App() {
  const { state, historyPreview, handleInput, handleExpressionChange, handleKeyDown } = useCalculator();

  return (
    <div className="app-shell">
      <div className="calculator-card">
        <p className="eyebrow">Prashant's calculator</p>
        <h1>Calculator</h1>
        <p className="subtitle">Expressions are validated locally before they are sent to the backend.</p>
        <CalculatorDisplay
          value={state.expression}
          onChange={handleExpressionChange}
          onKeyDown={handleKeyDown}
          result={state.result}
          loading={state.loading}
          error={state.error}
        />
        <Keypad buttons={CALCULATOR_BUTTONS} onButtonClick={handleInput} />
        <HistoryPanel entries={historyPreview} />
      </div>
    </div>
  );
}

export default App;
