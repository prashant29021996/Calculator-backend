export function validateExpression(expression: string): string | null {
  const trimmed = expression.trim();

  if (!trimmed) {
    return 'Expression is required';
  }

  if (trimmed.length > 256) {
    return 'Expression is too long';
  }

  const alphaTokens = trimmed.match(/[A-Za-z]+/g) ?? [];
  if (alphaTokens.some((token) => !/^sqrt$/i.test(token))) {
    return 'Expression contains invalid characters';
  }

  if (!/^[0-9+\-*/%^().\s]+$/.test(trimmed.replace(/\s+/g, '').replace(/sqrt/gi, ''))) {
    return 'Expression contains invalid characters';
  }

  if (trimmed.includes('/0') || trimmed.includes('%0')) {
    return 'Division by zero is not allowed';
  }

  if (/[+\-*/%^.](?=[+\-*/%^.])/.test(trimmed)) {
    return 'Expression is invalid';
  }

  if (/[+\-*/%^.]$/.test(trimmed)) {
    return 'Expression is invalid';
  }

  if (/^\./.test(trimmed) || /\.$/.test(trimmed)) {
    return 'Expression is invalid';
  }

  if ((trimmed.match(/\(/g) ?? []).length !== (trimmed.match(/\)/g) ?? []).length) {
    return 'Expression is invalid';
  }

  return null;
}
