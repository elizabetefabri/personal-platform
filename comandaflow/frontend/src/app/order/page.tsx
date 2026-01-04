"use client";

import { api } from "@/lib/api";
import { useEffect, useMemo, useState } from "react";

type MenuResponse = {
  status: "ok";
  data: {
    categories: string[];
    menu: Record<
      string,
      {
        id: number;
        name: string;
        price: number;
        image_url: string | null;
        type: string;
      }[]
    >;
  };
};

type CreateOrderResponse = {
  status: "ok";
  data: any;
};

export default function NewOrderPage() {
  const [client, setClient] = useState("Cliente Teste");
  const [menu, setMenu] = useState<MenuResponse["data"] | null>(null);
  const [selected, setSelected] = useState<Record<number, number>>({}); // productId -> qty
  const [msg, setMsg] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      const res = await api<MenuResponse>("/menu.php", {
        method: "GET",
        auth: true,
      });
      setMenu(res.data);
    })();
  }, []);

  const flatProducts = useMemo(() => {
    if (!menu) return [];
    return menu.categories.flatMap((c) => menu.menu[c]);
  }, [menu]);

  function inc(productId: number) {
    setSelected((s) => ({ ...s, [productId]: (s[productId] ?? 0) + 1 }));
  }

  function dec(productId: number) {
    setSelected((s) => {
      const next = { ...s };
      const v = (next[productId] ?? 0) - 1;
      if (v <= 0) delete next[productId];
      else next[productId] = v;
      return next;
    });
  }

  async function submit() {
    setMsg(null);
    const products = Object.entries(selected).map(([product_id, qty]) => ({
      product_id: Number(product_id),
      qty: Number(qty),
    }));

    if (!client.trim()) return setMsg("Client é obrigatório.");
    if (products.length === 0) return setMsg("Selecione ao menos 1 item.");

    try {
      const res = await api<CreateOrderResponse>("/orders.php", {
        method: "POST",
        auth: true,
        body: JSON.stringify({ client, products }),
      });
      setMsg("Pedido criado com sucesso.");
      setSelected({});
      console.log(res);
    } catch (err: any) {
      setMsg(err.message || "Erro ao criar pedido.");
    }
  }

  if (!menu) return <main style={{ padding: 24 }}>Carregando...</main>;

  return (
    <main style={{ padding: 24 }}>
      <h1>Novo Pedido</h1>

      <div style={{ display: "grid", gap: 12, maxWidth: 520 }}>
        <input
          value={client}
          onChange={(e) => setClient(e.target.value)}
          placeholder="Nome do cliente"
        />

        <h2>Produtos</h2>
        <ul style={{ display: "grid", gap: 8 }}>
          {flatProducts.map((p) => (
            <li
              key={p.id}
              style={{ display: "flex", gap: 12, alignItems: "center" }}
            >
              <span style={{ flex: 1 }}>
                {p.name} — R$ {(p.price / 100).toFixed(2)}
              </span>

              <button onClick={() => dec(p.id)}>-</button>
              <span>{selected[p.id] ?? 0}</span>
              <button onClick={() => inc(p.id)}>+</button>
            </li>
          ))}
        </ul>

        <button onClick={submit}>Criar Pedido</button>

        {msg && <p>{msg}</p>}
      </div>
    </main>
  );
}
