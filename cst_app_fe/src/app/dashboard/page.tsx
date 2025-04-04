'use client';

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import ClientLayout from "@/components/ClientLayout";
import api from "@/lib/api";

type Menu = {
  id: number;
  form_name: string;
  form_url: string;
};

export default function DashboardPage() {
  const [menus, setMenus] = useState<Menu[]>([]);
  const [error, setError] = useState('');
  const router = useRouter();

  useEffect(() => {
    const fetchMenus = async () => {
      const token = localStorage.getItem("token");
      try {
        const res = await api.get("/menu", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setMenus(res.data.data);
      } catch (err: any) {
        setError("Gagal ambil menu: " + (err.response?.data?.message || err.message));
      }
    };

    fetchMenus();
  }, []);

  return (
    <ClientLayout>
      <h1 className="text-xl font-bold mb-4">Daftar Menu</h1>
      {error && <p className="text-red-500">{error}</p>}
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {menus.map((menu) => (
          <div
            key={menu.id}
            onClick={() => router.push(menu.form_url)}
            className="cursor-pointer p-4 bg-white shadow rounded-xl"
          >
            <h2 className="font-semibold">{menu.form_name}</h2>
            <p className="text-sm text-gray-500">{menu.form_url}</p>
          </div>
        ))}
      </div>
    </ClientLayout>
  );
}
