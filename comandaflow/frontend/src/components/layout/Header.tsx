// src/components/layout/Header.tsx
"use client";

import { Suspense } from "react";
import HeaderContent from "./HeaderContent";

export default function Header() {
  return (
    <Suspense fallback={<header />}>
      <HeaderContent />
    </Suspense>
  );
}
