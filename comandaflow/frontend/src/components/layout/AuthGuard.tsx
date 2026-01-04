"use client";

import { authService } from "@/services/auth.service";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

const PUBLIC_ROUTES = ["/", "/login", "/cadastro"];

export default function AuthGuard({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const router = useRouter();
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    const logged = authService.isAuthenticated();
    const isPublic = PUBLIC_ROUTES.includes(pathname);

    // Se não estiver logado, força ir para /login (exceto se já estiver lá)
    if (!logged && !isPublic) {
      router.replace("/login");
      return;
    }

    // Se estiver logado e tentar acessar /login, manda para /menu
    if (logged && pathname === "/login") {
      router.replace("/menu");
      return;
    }

    setIsReady(true);
  }, [pathname, router]);

  if (!isReady) {
    return null; // Renderiza nada enquanto valida
  }

  return <>{children}</>;
}
