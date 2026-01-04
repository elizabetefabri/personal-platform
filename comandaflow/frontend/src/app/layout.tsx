"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import Footer from "../components/layout/Footer";
import Header from "../components/layout/Header";
import { authService } from "../services/auth.service";
import "../styles/globals.css";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();
  const router = useRouter();
  const [isAuth, setIsAuth] = useState(true);

  useEffect(() => {
    const logged = authService.isAuthenticated();
    setIsAuth(logged);

    // Se não estiver logado, força ir para /login (exceto se já estiver lá)
    if (!logged && pathname !== "/login") {
      router.replace("/login");
    }

    // Se estiver logado e tentar acessar /login, manda para /menu
    if (logged && pathname === "/login") {
      router.replace("/menu");
    }
  }, [pathname, router]);

  const showLayout = isAuth && pathname !== "/login";

  return (
    <html lang="pt-BR">
      <body>
        {showLayout && <Header />}
        {children}
        {showLayout && <Footer />}
      </body>
    </html>
  );
}
