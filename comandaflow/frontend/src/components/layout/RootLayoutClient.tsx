// src/components/layout/RootLayoutClient.tsx
"use client";

import { authService } from "@/services/auth.service";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useMemo, useState } from "react";
import { Toast } from "../ui/Toast";
import Footer from "./Footer";
import Header from "./Header";

const PUBLIC_PREFIXES = ["/", "/login", "/cadastro"];

function normalizePath(pathname: string) {
  if (!pathname) return "/";
  // remove "/" final (exceto se for "/")
  if (pathname.length > 1 && pathname.endsWith("/")) return pathname.slice(0, -1);
  return pathname;
}

export function RootLayoutClient({ children }: { children: React.ReactNode }) {
  const rawPathname = usePathname();
  const pathname = useMemo(() => normalizePath(rawPathname), [rawPathname]);

  const router = useRouter();
  const [isAuth, setIsAuth] = useState(false);
  const [isReady, setIsReady] = useState(false);

  const isPublic = useMemo(() => {
    // permite "/cadastro" e também futuros filhos como "/cadastro/..."
    return PUBLIC_PREFIXES.some((p) => (p === "/" ? pathname === "/" : pathname.startsWith(p)));
  }, [pathname]);

  useEffect(() => {
    const logged = authService.isAuthenticated();
    setIsAuth(logged);

    // se NÃO logado e rota NÃO pública -> /login
    if (!logged && !isPublic) {
      router.replace("/login");
      setIsReady(true);
      return;
    }

    // se logado e tentar /login ou /cadastro -> /menu
    if (logged && (pathname === "/login" || pathname === "/cadastro")) {
      router.replace("/menu");
      setIsReady(true);
      return;
    }

    setIsReady(true);
  }, [pathname, isPublic, router]);

  const showLayout = isReady && isAuth && !isPublic;

  return (
    <>
      <Toast />
      {showLayout && <Header />}
      {children}
      {showLayout && <Footer />}
    </>
  );
}
