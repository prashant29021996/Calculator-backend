import { render, screen } from '@testing-library/react';
import App from './App';

describe('App', () => {
  it('renders the calculator UI', () => {
    render(<App />);
    expect(screen.getByText('Calculator')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Enter an expression')).toBeInTheDocument();
  });
});
