const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080';

function snakeToCamel(str: string): string {
  return str.replace(/_([a-z0-9])/g, (_, c) => c.toUpperCase());
}

function camelToSnake(str: string): string {
  return str.replace(/[A-Z]/g, (c) => `_${c.toLowerCase()}`);
}

function transformKeys(obj: unknown, fn: (key: string) => string): unknown {
  if (Array.isArray(obj)) return obj.map((item) => transformKeys(item, fn));
  if (obj !== null && typeof obj === 'object') {
    return Object.fromEntries(
      Object.entries(obj as Record<string, unknown>).map(([k, v]) => [fn(k), transformKeys(v, fn)]),
    );
  }
  return obj;
}

export async function fetchAPI<T>(path: string, options?: RequestInit): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`, options);
  if (!res.ok) {
    const error = await res.json().catch(() => ({ error: { message: 'API Error' } }));
    throw new Error(error.error?.message || 'API Error');
  }
  const json = await res.json();
  return transformKeys(json.data, snakeToCamel) as T;
}

export async function mutateAPI<T>(
  path: string,
  method: 'POST' | 'PUT' | 'PATCH' | 'DELETE',
  body?: unknown,
): Promise<T> {
  const snakeBody = body ? transformKeys(body, camelToSnake) : undefined;
  return fetchAPI<T>(path, {
    method,
    headers: { 'Content-Type': 'application/json' },
    body: snakeBody ? JSON.stringify(snakeBody) : undefined,
  });
}
