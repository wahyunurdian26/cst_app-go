'use client';


import { ReactNode } from "react";

export default function ClientLayout({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen bg-gray-100">
      <header className="bg-white shadow p-4">
        <h1 className="text-xl font-bold">My App</h1>
      </header>

      <main className="p-6">{children}</main>

      <footer className="bg-white text-center p-4 mt-6 shadow-inner">
        <p className="text-sm text-gray-500">© 2025 My App</p>
      </footer>
    </div>
  );
}
