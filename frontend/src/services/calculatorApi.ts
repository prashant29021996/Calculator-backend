export type CalculateResponse = {
  result?: number;
  error?: string;
};

export async function calculateExpression(expression: string): Promise<CalculateResponse> {
  const response = await fetch('/api/calculate', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ expression }),
  });

  const data = (await response.json()) as CalculateResponse;
  if (!response.ok) {
    throw new Error(data.error ?? 'Request failed');
  }

  return data;
}
