type KeypadProps = {
  buttons: string[];
  onButtonClick: (button: string) => void;
};

function Keypad({ buttons, onButtonClick }: KeypadProps) {
  return (
    <div className="keypad" role="grid" aria-label="Calculator keypad">
      {buttons.map((button) => (
        <button
          key={button}
          className={`key ${button === '=' ? 'operator' : ''} ${button === 'DEL' || button === 'C' ? 'action' : ''}`}
          onClick={() => onButtonClick(button)}
          type="button"
          aria-label={button === '=' ? 'Calculate expression' : `Insert ${button}`}
        >
          {button}
        </button>
      ))}
    </div>
  );
}

export default Keypad;
