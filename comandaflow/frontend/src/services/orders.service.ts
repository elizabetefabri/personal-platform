// src/services/order.service.ts
import { authService } from "@/services/auth.service";
import { productService } from "@/services/products.service";
import type { Order, OrderStatus } from "@/types/order";

type ModalListener = (open: boolean) => void;

class OrderService {
  private modalOpen = false;
  private modalListeners = new Set<ModalListener>();

  private baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "";

  onModalChange(listener: ModalListener) {
    this.modalListeners.add(listener);
    listener(this.modalOpen);
    return () => this.modalListeners.delete(listener);
  }

  openModal() {
    this.modalOpen = true;
    this.modalListeners.forEach((l) => l(true));
  }

  closeModal() {
    this.modalOpen = false;
    this.modalListeners.forEach((l) => l(false));
  }

  private headersJson() {
    const token = authService.getToken();
    return {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    };
  }

  async createOrder(payload: Omit<Order, "id">) {
    const res = await fetch(`${this.baseUrl}/orders.php`, {
      method: "POST",
      headers: this.headersJson(),
      body: JSON.stringify(payload),
    });

    return res.json();
  }

  async listOrdersByStatus(status: OrderStatus) {
    const res = await fetch(
      `${this.baseUrl}/orders.php?status=${encodeURIComponent(status)}`,
      {
        method: "GET",
        headers: this.headersJson(),
      }
    );

    return res.json();
  }

  async updateOrderStatus(orderId: number, status: OrderStatus) {
    const res = await fetch(`${this.baseUrl}/orders.php`, {
      method: "PATCH",
      headers: this.headersJson(),
      body: JSON.stringify({ id: orderId, status }),
    });

    return res.json();
  }

  /**
   * Limpa o carrinho no cliente e reseta o contador de itens.
   * Atualmente não há armazenamento local, então apenas zera o contador.
   */
  async clearCart() {
    productService.updateItemCount(0);
    return Promise.resolve();
  }
}

export const orderService = new OrderService();
