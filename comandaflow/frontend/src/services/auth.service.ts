import type { User } from "../types/user";

const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL;

if (!API_BASE) {
  throw new Error("NEXT_PUBLIC_API_BASE_URL não definido no .env.local");
}

type LoginResponse = {
  status: "ok";
  data: {
    token: string;
    user: { id: number; email: string; role: "admin" | "waiter" | "chef" };
  };
};

type ApiErrorResponse = {
  status?: "error";
  message?: string;
  error?: string;
};

const TOKEN_KEY = "cf_token";
const USER_KEY = "cf_user";

function isBrowser(): boolean {
  return typeof window !== "undefined";
}

function getStoredToken(): string | null {
  if (!isBrowser()) return null;
  return localStorage.getItem(TOKEN_KEY);
}

function setStoredAuth(token: string, user: User) {
  if (!isBrowser()) return;
  localStorage.setItem(TOKEN_KEY, token);
  localStorage.setItem(USER_KEY, JSON.stringify(user));
}

function clearStoredAuth() {
  if (!isBrowser()) return;
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(USER_KEY);
}

function buildAuthHeaders(token: string | null): HeadersInit {
  if (!token) return {};
  return { Authorization: `Bearer ${token}` };
}

async function parseError(res: Response): Promise<string> {
  const data = (await res.json().catch(() => null)) as ApiErrorResponse | null;
  return data?.message || data?.error || `Erro HTTP ${res.status}`;
}

export const authService = {
  // ✅ para outros services chamarem (Order/Product/etc.)
  getToken(): string | null {
    return getStoredToken();
  },

  // ✅ útil para reaproveitar headers sem copiar regra
  getAuthHeaders(): HeadersInit {
    return buildAuthHeaders(getStoredToken());
  },

  // ✅ opcional: força erro se não estiver logado (bom para POST/PUT/DELETE)
  requireAuthHeaders(): HeadersInit {
    const token = getStoredToken();
    if (!token) throw new Error("Não autenticado: token ausente.");
    return buildAuthHeaders(token);
  },

  isAuthenticated(): boolean {
    return !!getStoredToken();
  },

  getCurrentUser(): User | null {
    if (!isBrowser()) return null;
    const raw = localStorage.getItem(USER_KEY);
    if (!raw) return null;
    try {
      return JSON.parse(raw) as User;
    } catch {
      return null;
    }
  },

  async login(
    email: string,
    password: string
  ): Promise<{ token: string; user: User }> {
    const res = await fetch(`${API_BASE}/auth_login.php`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });

    if (!res.ok) {
      throw new Error(await parseError(res));
    }

    const json = (await res.json()) as LoginResponse;
    const token = json.data.token;

    const user: User = {
      id: String(json.data.user.id),
      name: "",
      email: json.data.user.email,
      role: json.data.user.role,
    };

    setStoredAuth(token, user);
    return { token, user };
  },

  logout(): void {
    clearStoredAuth();
  },

  // CRUD Users (no seu backend: users.php)
  async registerUser(user: User): Promise<User> {
    const res = await fetch(`${API_BASE}/users.php`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...authService.requireAuthHeaders(),
      },
      body: JSON.stringify(user),
    });

    if (!res.ok) throw new Error(await parseError(res));

    const data = (await res.json()) as any;
    return (data?.data ?? data) as User;
  },

  async updateUser(id: string, user: User): Promise<User> {
    const res = await fetch(
      `${API_BASE}/users.php?id=${encodeURIComponent(id)}`,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          ...authService.requireAuthHeaders(),
        },
        body: JSON.stringify(user),
      }
    );

    if (!res.ok) throw new Error(await parseError(res));

    const data = (await res.json()) as any;
    return (data?.data ?? data) as User;
  },

  async deleteUser(id: string): Promise<void> {
    const res = await fetch(
      `${API_BASE}/users.php?id=${encodeURIComponent(id)}`,
      {
        method: "DELETE",
        headers: {
          ...authService.requireAuthHeaders(),
        },
      }
    );

    if (!res.ok) throw new Error(await parseError(res));
  },
};
