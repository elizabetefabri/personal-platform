export interface User {
  id?: string;
  name: string;
  email: string;
  role: "admin" | "waiter" | "chef" | string;
  password?: string;
}
