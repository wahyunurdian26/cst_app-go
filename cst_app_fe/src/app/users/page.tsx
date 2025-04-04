'use client';

import { useEffect, useState } from "react";
import api from "@/lib/api";
import ClientLayout from "@/components/ClientLayout";

import { useRouter } from "next/navigation";
import Layout from "@/components/layouts";


type User = {
  id: string;
  email: string;
  username: string;
  id_role: string;
  id_business_group: string;
  id_sub_business_group: string;
  email_pic: string;
  status_active: boolean;
  id_business_group_digital: string;
  created_at: string;
  updated_at: string;
};

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [error, setError] = useState("");
  const router = useRouter();

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    const token = localStorage.getItem("token");
    try {
      const res = await api.get("/users", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setUsers(res.data.data);
    } catch (err: any) {
      setError("Gagal mengambil data: " + (err.response?.data?.message || err.message));
    }
  };

  const handleDelete = async (id: string) => {
    if (!confirm("Apakah Anda yakin ingin menghapus user ini?")) return;

    const token = localStorage.getItem("token");
    try {
      await api.delete(`/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      alert("User berhasil dihapus!");
      fetchUsers();
    } catch (err: any) {
      alert("Gagal menghapus user: " + (err.response?.data?.message || err.message));
    }
  };

  return (
    <Layout>
      <h1 className="text-xl font-bold text-gray-800 mb-4">Daftar Pengguna</h1>
      <button
        onClick={() => router.push("/users/create")}
        className="mb-4 bg-blue-500 text-white px-4 py-2 rounded-md"
      >
        + Tambah User
      </button>
      {error && <p className="text-red-500">{error}</p>}
      <div className="overflow-x-auto">
        <table className="min-w-full border border-gray-300 bg-white">
          <thead>
            <tr className="bg-gray-200">
              <th className="border px-4 py-2">ID</th>
              <th className="border px-4 py-2">Email</th>
              <th className="border px-4 py-2">Username</th>
              <th className="border px-4 py-2">Role</th>
              <th className="border px-4 py-2">Business Group</th>
              <th className="border px-4 py-2">Status</th>
              <th className="border px-4 py-2">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.id} className="hover:bg-gray-100">
                <td className="border px-4 py-2">{user.id}</td>
                <td className="border px-4 py-2">{user.email}</td>
                <td className="border px-4 py-2">{user.username}</td>
                <td className="border px-4 py-2">{user.id_role}</td>
                <td className="border px-4 py-2">{user.id_business_group}</td>
                <td className={`border px-4 py-2 font-bold ${user.status_active ? "text-green-600" : "text-red-600"}`}>
                  {user.status_active ? "Aktif" : "Tidak Aktif"}
                </td>
                <td className="border px-4 py-2">
                  <button
                    onClick={() => router.push(`/users/edit/${user.id}`)}
                    className="bg-yellow-500 text-white px-3 py-1 rounded mr-2"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDelete(user.id)}
                    className="bg-red-500 text-white px-3 py-1 rounded"
                  >
                    Hapus
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
}
