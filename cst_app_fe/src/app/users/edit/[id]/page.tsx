'use client';

import { useParams, useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';

interface User {
  id: string;
  email: string;
  username: string;
  password: string;
  id_role: string;
  id_business_group: string;
  id_sub_business_group: string;
  email_pic: string;
  status_active: boolean;
  id_business_group_digital: string;
  created_at: string;
  updated_at: string;
}

export default function EditUserPage() {
  const router = useRouter();
  const params = useParams<{ id: string }>();
  const id = params?.id;

  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    if (!id) return;

    const token = localStorage.getItem('token');
    fetch(`http://localhost:8080/api/users/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then(res => res.json())
      .then(data => setUser(data.data));
  }, [id]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setUser(prev => prev ? { ...prev, [name]: name === "status_active" ? value === "true" : value } : null);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem('token');
  
    console.log('Sending data:', user);
  
    const res = await fetch(`http://localhost:8080/api/users/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(user),
    });
  
    const text = await res.text();
    console.log('Response:', text);
  
    if (res.ok) {
      alert('User berhasil diupdate!');
      router.push('/users');
    } else {
      alert('Gagal update user.');
      console.error('Gagal:', text);
    }
  };
  
  if (!user) return <div>Loading...</div>;

  return (
    <div className="p-6 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Edit User: {user.username}</h1>
      <form onSubmit={handleSubmit} className="flex flex-col gap-3">
        <input name="email" value={user.email} onChange={handleChange} className="border px-3 py-2" />
        <input name="username" value={user.username} onChange={handleChange} className="border px-3 py-2" />
        <input name="password" value={user.password} onChange={handleChange} className="border px-3 py-2" />
        <input name="id_role" value={user.id_role} onChange={handleChange} className="border px-3 py-2" />
        <input name="id_business_group" value={user.id_business_group} onChange={handleChange} className="border px-3 py-2" />
        <input name="id_sub_business_group" value={user.id_sub_business_group} onChange={handleChange} className="border px-3 py-2" />
        <input name="email_pic" value={user.email_pic} onChange={handleChange} className="border px-3 py-2" />
        <select
          name="status_active"
          value={user.status_active ? 'true' : 'false'}
          onChange={handleChange}
          className="border px-3 py-2"
        >
          <option value="true">Active</option>
          <option value="false">Inactive</option>
        </select>
        <input name="id_business_group_digital" value={user.id_business_group_digital} onChange={handleChange} className="border px-3 py-2" />
        <button className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
          Simpan
        </button>
      </form>
    </div>
  );
}
