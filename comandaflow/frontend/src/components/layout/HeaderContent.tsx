// src/components/layout/HeaderContent.tsx
"use client";

import { useRouter, useSearchParams } from "next/navigation";
import { useEffect, useMemo, useState } from "react";

import styles from "@/styles/components/layout/header.module.css";

import { FiArrowLeft, FiClipboard, FiLogOut, FiMenu } from "react-icons/fi";

import { authService } from "@/services/auth.service";
import { orderService } from "@/services/orders.service";
import { productService } from "@/services/products.service";

export default function HeaderContent() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const mesaId = useMemo(
    () => searchParams.get("mesaId") ?? "",
    [searchParams]
  );

  const [itemCount, setItemCount] = useState(0);
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  useEffect(() => {
    const unsubscribe = productService.onItemCountChange((count) => {
      setItemCount(count);
    });
    return () => {
      if (typeof unsubscribe === "function") {
        unsubscribe();
      }
    };
  }, []);

  function goToMesa() {
    const url = mesaId ? `/menu?mesaId=${encodeURIComponent(mesaId)}` : "/menu";
    router.push(url);
  }

  function toggleMenu() {
    setIsMenuOpen((prev) => !prev);
  }

  function closeMenu() {
    setIsMenuOpen(false);
  }

  async function handleLogout() {
    authService.logout();
    await orderService.clearCart();
    router.replace("/login");
  }

  return (
    <header className={styles.header}>
      <div className={styles.container}>
        <button
          className={styles.backButton}
          onClick={goToMesa}
          aria-label="Voltar para mesa"
        >
          <FiArrowLeft size={20} />
        </button>

        <h1 className={styles.title}>ComandaFlow</h1>

        <nav className={styles.nav}>
          <button
            className={styles.cartButton}
            onClick={() => router.push("/order")}
            aria-label={`Carrinho com ${itemCount} itens`}
          >
            <FiClipboard size={20} />
            {itemCount > 0 && <span className={styles.badge}>{itemCount}</span>}
          </button>

          <button
            className={styles.menuButton}
            onClick={toggleMenu}
            aria-label="Menu"
          >
            <FiMenu size={20} />
          </button>
        </nav>
      </div>

      {isMenuOpen && (
        <div
          className={styles.backdrop}
          onClick={closeMenu}
          role="presentation"
        />
      )}

      {isMenuOpen && (
        <ul className={styles.menu}>
          <li>
            <button onClick={handleLogout}>
              <FiLogOut size={18} />
              Sair
            </button>
          </li>
        </ul>
      )}
    </header>
  );
}
