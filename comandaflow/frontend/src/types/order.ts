import type { Product } from "./product";

export type OrderStatus = "pending" | "preparing" | "done" | "canceled";

export interface OrderItem {
  product: Product;
  quantity: number;
  unit_price: number; // em centavos
}

export interface Order {
  id?: number;
  user_id: number;
  table?: string | null;
  status: OrderStatus;
  items: OrderItem[];
  created_at?: string;
  updated_at?: string | null;
}
