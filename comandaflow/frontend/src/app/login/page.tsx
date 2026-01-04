"use client";

import { authService } from "@/services/auth.service";
import styles from "@/styles/pages/login.module.css";
import { useRouter } from "next/navigation";
import { useMemo, useState } from "react";
import { FiEye, FiEyeOff } from "react-icons/fi";

export default function LoginPage() {
  const router = useRouter();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    setLoading(true);

    try {
      // Autentica no backend (MySQL) via auth_login.php e salva token/user no localStorage (authService)
      await authService.login(email.trim(), password);
      router.replace("/menu");
    } catch (err: any) {
      setError(err?.message ?? "Email ou senha inválidos.");
    } finally {
      setLoading(false);
    }
  }

  const canSubmit = useMemo(() => {
    return email.trim().length > 0 && password.length >= 6 && !loading;
  }, [email, password, loading]);

  return (
    <section className={styles.container}>
      <main className={styles.card}>
        <div className={styles.logoWrapper}>
          <img
            className={styles.logo}
            src="/images/logo.png"
            alt="Logo do ComandaFlow"
          />
        </div>

        <div className={styles.formWrapper}>
          <h3 className={styles.title}>Login</h3>

          <form onSubmit={onSubmit} className={styles.form}>
            <input
              className={styles.input}
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              type="email"
              placeholder="Digite seu e-mail"
              id="email"
              autoComplete="username"
            />

            <div className={styles.passwordField}>
              <input
                className={styles.input}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type={showPassword ? "text" : "password"}
                placeholder="Senha"
                id="password"
                autoComplete="current-password"
              />

              <button
                type="button"
                className={styles.iconButton}
                onClick={() => setShowPassword((v) => !v)}
                aria-label={showPassword ? "Ocultar senha" : "Mostrar senha"}
              >
                {showPassword ? <FiEyeOff size={18} /> : <FiEye size={18} />}
              </button>
            </div>

            <button
              type="submit"
              className={styles.primaryButton}
              disabled={!canSubmit}
            >
              {loading ? "Entrando..." : "Entrar"}
            </button>

            <div className={styles.registerRow}>
              <span className={styles.registerText}>Não possui cadastro?</span>
              <button
                type="button"
                className={styles.secondaryButton}
                onClick={() => router.push("/cadastro")}
              >
                Cadastre-se
              </button>
            </div>

            {error && <p className={styles.error}>{error}</p>}
          </form>
        </div>

        <footer className={styles.footer}>
          <p className={styles.footerText}>
            ComandaFlow© 2026{" "}
            <em className={styles.footerEm}>
              <a
                className={styles.footerLink}
                href="https://elizabetefsousafabri.com.br"
                target="_blank"
                rel="noreferrer"
              >
                Desenvolvido por Elizabete Fabri
              </a>
            </em>
          </p>
        </footer>
      </main>
    </section>
  );
}
