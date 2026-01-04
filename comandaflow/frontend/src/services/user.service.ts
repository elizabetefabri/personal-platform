import { authService } from "./auth.service";

export interface User {
  id?: string;
  name: string;
  email: string;
  password: string;
  role: "admin" | "waiter" | "chef";
}

class UserService {
  private baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "";

  private getAuthorizationHeader() {
    const token = authService.getToken();
    return {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    };
  }

  async registerUser(user: Omit<User, "id">): Promise<User> {
    try {
      const response = await fetch(`${this.baseUrl}/users`, {
        method: "POST",
        headers: this.getAuthorizationHeader(),
        body: JSON.stringify(user),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Erro ao cadastrar usuário. 😕");
      }

      return response.json();
    } catch (error: any) {
      throw new Error(error.message || "Erro ao cadastrar usuário. 😕");
    }
  }

  async listUsers(): Promise<User[]> {
    try {
      const response = await fetch(`${this.baseUrl}/users`, {
        method: "GET",
        headers: this.getAuthorizationHeader(),
      });

      if (!response.ok) {
        throw new Error("Erro ao carregar usuários. 😕");
      }

      return response.json();
    } catch (error) {
      console.error("Error listing users", error);
      return [];
    }
  }

  async deleteUser(id: string): Promise<void> {
    try {
      const response = await fetch(`${this.baseUrl}/users/${id}`, {
        method: "DELETE",
        headers: this.getAuthorizationHeader(),
      });

      if (!response.ok) {
        throw new Error("Erro ao excluir usuário. 😕");
      }
    } catch (error) {
      console.error("Error deleting user", error);
      throw error;
    }
  }

  async updateUser(user: User): Promise<User> {
    try {
      const response = await fetch(`${this.baseUrl}/users/${user.id}`, {
        method: "PUT",
        headers: this.getAuthorizationHeader(),
        body: JSON.stringify(user),
      });

      if (!response.ok) {
        throw new Error("Erro ao atualizar usuário. 😕");
      }

      return response.json();
    } catch (error: any) {
      throw new Error(error.message || "Erro ao atualizar usuário. 😕");
    }
  }
}

export const userService = new UserService();
