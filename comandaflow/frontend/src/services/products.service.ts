import { authService } from "@/services/auth.service";
import type { Product } from "@/types/product";

type ItemCountListener = (count: number) => void;

class ProductService {
  private baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "";

  private itemCount = 0;
  private listeners = new Set<ItemCountListener>();

  onItemCountChange(listener: ItemCountListener) {
    this.listeners.add(listener);
    listener(this.itemCount);
    return () => this.listeners.delete(listener);
  }

  updateItemCount(count: number) {
    this.itemCount = count;
    this.listeners.forEach((l) => l(count));
  }

  private headersJson() {
    const token = authService.getToken();
    return {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    };
  }

  async createProduct(product: Omit<Product, "id">) {
    const res = await fetch(`${this.baseUrl}/products.php`, {
      method: "POST",
      headers: this.headersJson(),
      body: JSON.stringify(product),
    });
    return res.json();
  }

  async listProducts() {
    const res = await fetch(`${this.baseUrl}/products.php`, {
      method: "GET",
      headers: this.headersJson(),
    });
    return res.json();
  }

  async getProductById(id: number) {
    const res = await fetch(`${this.baseUrl}/products.php?id=${id}`, {
      method: "GET",
      headers: this.headersJson(),
    });
    return res.json();
  }

  async updateProduct(product: Product) {
    const res = await fetch(`${this.baseUrl}/products.php`, {
      method: "PUT",
      headers: this.headersJson(),
      body: JSON.stringify(product),
    });
    return res.json();
  }

  async deleteProduct(id: number) {
    const res = await fetch(`${this.baseUrl}/products.php?id=${id}`, {
      method: "DELETE",
      headers: this.headersJson(),
    });
    return res.json();
  }
}

export const productService = new ProductService();
