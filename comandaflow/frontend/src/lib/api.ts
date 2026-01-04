const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL;

if (!API_BASE) {
  throw new Error("NEXT_PUBLIC_API_BASE_URL não definido no .env.local");
}

export function getToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("cf_token");
}

export function setToken(token: string) {
  localStorage.setItem("cf_token", token);
}

export function clearToken() {
  localStorage.removeItem("cf_token");
}

type ApiOptions = RequestInit & { auth?: boolean };

export async function api<T>(path: string, options: ApiOptions = {}): Promise<T> {
  const headers = new Headers(options.headers);

  // JSON padrão
  if (!headers.has("Content-Type") && options.body) {
    headers.set("Content-Type", "application/json");
  }

  // Auth
  if (options.auth) {
    const token = getToken();
    if (!token) throw new Error("Sem token (usuário não autenticado).");
    headers.set("Authorization", `Bearer ${token}`);
  }

  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers,
  });

  const data = await res.json().catch(() => null);

  if (!res.ok) {
    const msg = data?.message || `Erro HTTP ${res.status}`;
    throw new Error(msg);
  }

  return data as T;
}
