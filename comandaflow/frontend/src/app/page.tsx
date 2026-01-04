// src/app/page.tsx
"use client";

import { authService } from "@/services/auth.service";
import styles from "@/styles/pages/login.module.css";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    const isAuth = authService.isAuthenticated();
    router.replace(isAuth ? "/menu" : "/login");
  }, [router]);

  return <div className={styles.page}>Carregando...</div>;
}
