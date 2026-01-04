"use client";

import { useEffect, useState } from "react";
import { api } from "@/lib/api";

type MenuResponse = {
  status: "ok";
  data: {
    categories: string[];
    menu: Record<
      string,
      { id: number; name: string; price: number; image_url: string | null; type: string }[]
    >;
  };
};

export default function MenuPage() {
  const [data, setData] = useState<MenuResponse["data"] | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const res = await api<MenuResponse>("/menu.php", { method: "GET", auth: true });
        setData(res.data);
      } catch (err: any) {
        setError(err.message || "Erro ao carregar menu.");
      }
    })();
  }, []);

  if (error) return <main style={{ padding: 24 }}>{error}</main>;
  if (!data) return <main style={{ padding: 24 }}>Carregando menu...</main>;

  return (
    <main style={{ padding: 24 }}>
      <h1>Menu</h1>

      {data.categories.map((cat) => (
        <section key={cat} style={{ marginTop: 20 }}>
          <h2>{cat}</h2>
          <ul>
            {data.menu[cat].map((p) => (
              <li key={p.id}>
                #{p.id} — {p.name} — R$ {(p.price / 100).toFixed(2)}
              </li>
            ))}
          </ul>
        </section>
      ))}
    </main>
  );
}
