export type ProductType = "breakfast" | "lunch" | "dinner" | "drink" | "other";

export interface Product {
  id?: number;
  name: string;
  price: number; // em centavos
  image_url?: string | null;
  type: ProductType;
  is_active: boolean;
  created_at?: string;
  updated_at?: string | null;
}
