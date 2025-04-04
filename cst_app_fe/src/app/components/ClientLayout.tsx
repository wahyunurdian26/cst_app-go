'use client';

import { ReactNode } from "react";

export default function ClientLayout({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen flex flex-col bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow px-6 py-4">
        <h1 className="text-xl font-semibold text-gray-800">CST App</h1>
      </header>

      {/* Main content */}
      <main className="flex-1 px-6 py-4">
        {children}
      </main>

      {/* Footer */}
      <footer className="bg-white text-center py-4 border-t">
        <p className="text-sm text-gray-500">© 2025 CST App</p>
      </footer>
    </div>
  );
}
