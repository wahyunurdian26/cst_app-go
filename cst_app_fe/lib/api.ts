// lib/api.ts
import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080/api",
});

// Tambahkan token secara dinamis
api.interceptors.request.use(
  (config) => {
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("token");
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export default api;
