// hooks/useMenu.ts
import { useEffect, useState } from "react";
import api from "@/lib/api";

export function useMenu() {
  const [menus, setMenus] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchMenu = async () => {
      try {
        const res = await api.get("/menu");
        setMenus(res.data.data);
      } catch (err: any) {
        setError("Gagal ambil menu: " + (err.response?.data?.message || ""));
      } finally {
        setLoading(false);
      }
    };

    fetchMenu();
  }, []);

  return { menus, loading, error };
}
