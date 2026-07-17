type HistoryPanelProps = {
  entries: string[];
};

function HistoryPanel({ entries }: HistoryPanelProps) {
  return (
    <div className="history-panel" aria-label="Recent calculations">
      <h2>Recent calculations</h2>
      {entries.length ? (
        <ul>
          {entries.map((entry) => (
            <li key={entry}>{entry}</li>
          ))}
        </ul>
      ) : (
        <p>No history yet.</p>
      )}
    </div>
  );
}

export default HistoryPanel;
