"use client";

import { useEffect, useState } from "react";
import styles from "@/styles/components/ui/toast.module.css";

export type ToastType = "success" | "error" | "info";

export interface ToastMessage {
  id: string;
  message: string;
  type: ToastType;
}

let toastId = 0;
const listeners = new Set<(toast: ToastMessage) => void>();

export function useToast() {
  const [toasts, setToasts] = useState<ToastMessage[]>([]);

  useEffect(() => {
    const handleAddToast = (toast: ToastMessage) => {
      setToasts((prev) => [...prev, toast]);

      // Auto-remove após 3 segundos
      setTimeout(() => {
        removeToast(toast.id);
      }, 3000);
    };

    listeners.add(handleAddToast);
    return () => {
      listeners.delete(handleAddToast);
    };
  }, []);

  return {
    toasts,
    show: (message: string, type: ToastType = "info") => {
      const id = `toast-${++toastId}`;
      const toast: ToastMessage = { id, message, type };
      listeners.forEach((listener) => listener(toast));
    },
  };
}

export function removeToast(id: string) {
  // Notificar todos os listeners para remover
  listeners.forEach((listener) =>
    listener({ id, message: "", type: "info" })
  );
}

export function Toast() {
  const { toasts } = useToast();

  const [displayToasts, setDisplayToasts] = useState<ToastMessage[]>([]);

  useEffect(() => {
    if (toasts.length > 0) {
      const lastToast = toasts[toasts.length - 1];
      if (lastToast.message) {
        setDisplayToasts((prev) => [...prev, lastToast]);
      } else {
        // Remove toast
        setDisplayToasts((prev) => prev.filter((t) => t.id !== lastToast.id));
      }
    }
  }, [toasts]);

  return (
    <div className={styles.container}>
      {displayToasts.map((toast) => (
        <div key={toast.id} className={`${styles.toast} ${styles[toast.type]}`}>
          {toast.message}
        </div>
      ))}
    </div>
  );
}
