import { fireEvent, render, screen } from '@testing-library/react';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import App from './App';

describe('App', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it('renders the calculator UI', () => {
    render(<App />);
    expect(screen.getByText('Calculator')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Enter an expression')).toBeInTheDocument();
  });

  it('validates expressions before sending them to the backend', async () => {
    const fetchMock = vi.fn();
    vi.stubGlobal('fetch', fetchMock);

    render(<App />);
    const input = screen.getByRole('textbox', { name: /expression/i });

    fireEvent.change(input, { target: { value: '2+' } });
    fireEvent.click(screen.getByRole('button', { name: /calculate expression/i }));

    expect(fetchMock).not.toHaveBeenCalled();
    expect(await screen.findByText(/expression is invalid/i)).toBeInTheDocument();
  });

  it('sends a calculation request only when the user evaluates an expression', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 4 }),
    });
    vi.stubGlobal('fetch', fetchMock);

    render(<App />);
    const input = screen.getByRole('textbox', { name: /expression/i });

    fireEvent.change(input, { target: { value: '2+2' } });
    fireEvent.click(screen.getByRole('button', { name: /calculate expression/i }));

    expect(fetchMock).toHaveBeenCalledTimes(1);
    expect(await screen.findByText('4')).toBeInTheDocument();
  });

  it('filters out random text before it can be sent to the backend', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 4 }),
    });
    vi.stubGlobal('fetch', fetchMock);

    render(<App />);
    const input = screen.getByRole('textbox', { name: /expression/i });

    fireEvent.change(input, { target: { value: '2+abc' } });
    fireEvent.click(screen.getByRole('button', { name: /calculate expression/i }));

    expect(fetchMock).not.toHaveBeenCalled();
    expect(await screen.findByText(/expression contains invalid characters/i)).toBeInTheDocument();
  });

  it('supports percent and exponentiation buttons', () => {
    render(<App />);
    const input = screen.getByRole('textbox', { name: /expression/i });

    fireEvent.click(screen.getByRole('button', { name: /insert %/i }));
    fireEvent.click(screen.getByRole('button', { name: /insert \^/i }));

    expect(input).toHaveValue('%^');
  });

  it('converts exponentiation to the backend-safe form before sending the request', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 25 }),
    });
    vi.stubGlobal('fetch', fetchMock);

    render(<App />);
    const input = screen.getByRole('textbox', { name: /expression/i });

    fireEvent.change(input, { target: { value: '5^2' } });
    fireEvent.click(screen.getByRole('button', { name: /calculate expression/i }));

    expect(fetchMock).toHaveBeenCalledWith(
      'http://localhost:8081/calculate',
      expect.objectContaining({
        body: JSON.stringify({ expression: '5**2' }),
      }),
    );
  });

  it('converts percentage expressions to a backend-safe form', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 7.5 }),
    });
    vi.stubGlobal('fetch', fetchMock);

    render(<App />);
    const input = screen.getByRole('textbox', { name: /expression/i });

    fireEvent.change(input, { target: { value: '25%30' } });
    fireEvent.click(screen.getByRole('button', { name: /calculate expression/i }));

    expect(fetchMock).toHaveBeenCalledWith(
      'http://localhost:8081/calculate',
      expect.objectContaining({
        body: JSON.stringify({ expression: '30*(25/100)' }),
      }),
    );
  });
});
