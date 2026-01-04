"use client";

import { SelectCustom } from "@/components/ui/SelectCustom";
import { useToast } from "@/components/ui/Toast";
import { authService } from "@/services/auth.service";
import { userService } from "@/services/user.service";
import styles from "@/styles/pages/cadastro.module.css";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function RegisterPage() {
  const router = useRouter();
  const { show: showToast } = useToast();

  const [formData, setFormData] = useState({
    name: "",
    email: "",
    confirmEmail: "",
    password: "",
    role: "",
  });

  const [loading, setLoading] = useState(false);
  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  const validateForm = () => {
    const newErrors: { [key: string]: string } = {};

    if (!formData.name.trim()) {
      newErrors.name = "Nome é obrigatório";
    }

    if (!formData.email.trim()) {
      newErrors.email = "Email é obrigatório";
    } else if (!isValidEmail(formData.email)) {
      newErrors.email = "Email inválido";
    }

    if (!formData.confirmEmail.trim()) {
      newErrors.confirmEmail = "Confirmação de email é obrigatória";
    } else if (formData.email !== formData.confirmEmail) {
      newErrors.confirmEmail = "Os emails não coincidem";
    }

    if (!formData.password.trim()) {
      newErrors.password = "Senha é obrigatória";
    } else if (formData.password.length < 6) {
      newErrors.password = "Senha deve ter no mínimo 6 caracteres";
    }

    if (!formData.role) {
      newErrors.role = "Cargo é obrigatório";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const isValidEmail = (email: string) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
    // Limpar erro ao começar a digitar
    if (errors[name]) {
      setErrors((prev) => ({
        ...prev,
        [name]: "",
      }));
    }
  };

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateForm()) {
      showToast("Por favor, preencha todos os campos corretamente", "error");
      return;
    }

    setLoading(true);

    try {
      const { name, email, password, role } = formData;

      await userService.registerUser({
        name,
        email,
        password,
        role: role as "admin" | "waiter" | "chef",
      });

      showToast("Usuário cadastrado com sucesso!", "success");

      // Fazer login automático com as credenciais do novo usuário
      setTimeout(async () => {
        try {
          await authService.login(email, password);
          router.replace("/menu");
        } catch (loginError) {
          // Se o login automático falhar, redirecionar para login manual
          console.error("Auto-login falhou", loginError);
          router.replace("/login");
        }
      }, 1500);
    } catch (error: any) {
      const errorMessage =
        error.message || "Erro ao cadastrar usuário. Tente novamente.";
      showToast(errorMessage, "error");
      console.error("Error registering user", error);
    } finally {
      setLoading(false);
    }
  };

  const canSubmit =
    formData.name.trim().length > 0 &&
    formData.email.trim().length > 0 &&
    formData.confirmEmail.trim().length > 0 &&
    formData.password.length >= 6 &&
    formData.role.length > 0 &&
    !loading;

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
          <h3 className={styles.title}>Cadastre-se</h3>

          <form onSubmit={onSubmit} className={styles.form}>
            <input
              className={styles.input}
              type="text"
              name="name"
              placeholder="Digite seu nome completo"
              value={formData.name}
              onChange={handleInputChange}
              disabled={loading}
            />
            {errors.name && <p className={styles.error}>{errors.name}</p>}

            <input
              className={styles.input}
              type="email"
              name="email"
              placeholder="Digite seu e-mail"
              value={formData.email}
              onChange={handleInputChange}
              disabled={loading}
              autoComplete="email"
            />
            {errors.email && <p className={styles.error}>{errors.email}</p>}

            <input
              className={styles.input}
              type="email"
              name="confirmEmail"
              placeholder="Confirme seu e-mail"
              value={formData.confirmEmail}
              onChange={handleInputChange}
              disabled={loading}
              autoComplete="email"
            />
            {errors.confirmEmail && (
              <p className={styles.error}>{errors.confirmEmail}</p>
            )}

            <input
              className={styles.input}
              type="password"
              name="password"
              placeholder="Senha"
              value={formData.password}
              onChange={handleInputChange}
              disabled={loading}
              autoComplete="new-password"
            />
            {errors.password && (
              <p className={styles.error}>{errors.password}</p>
            )}

            <div className={styles.roleRow}>
              <SelectCustom
                name="role"
                value={formData.role}
                onChange={handleInputChange}
                disabled={loading}
                options={[
                  { value: "", label: "Cargo" },
                  { value: "admin", label: "Admin" },
                  { value: "waiter", label: "Garçonete" },
                  { value: "chef", label: "Chefe de cozinha" },
                ]}
                ariaLabel="Selecione sua Função"
              />

              <button
                type="submit"
                className={styles.button}
                disabled={!canSubmit}
              >
                {loading ? "Cadastrando..." : "Cadastrar"}
              </button>
            </div>
            {errors.role && <p className={styles.error}>{errors.role}</p>}
          </form>

          <a href="/" className={styles.backLink}>
            ← Voltar
          </a>
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
