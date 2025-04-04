import { ReactNode } from "react";
import { useRouter } from "next/router";

type Props = {
  children: ReactNode;
};

export default function Layout({ children }: Props) {
  const router = useRouter();

  const handleLogout = () => {
    localStorage.removeItem("token");
    router.push("/login");
  };

  return (
    <div style={{ display: "flex", flexDirection: "column", minHeight: "100vh", fontFamily: "sans-serif" }}>
      {/* Header */}
      <header style={{ backgroundColor: "#f0f0f0", padding: "1rem", borderBottom: "1px solid #ddd" }}>
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
          <h3 style={{ margin: 0 }}>Dashboard</h3>
          <button
            onClick={handleLogout}
            style={{
              background: "#e74c3c",
              color: "#fff",
              border: "none",
              padding: "0.5rem 1rem",
              borderRadius: "4px",
              cursor: "pointer",
            }}
          >
            Logout
          </button>
        </div>
      </header>

      {/* Main content */}
      <main style={{ flex: 1, padding: "1.5rem", backgroundColor: "#fff" }}>
        {children}
      </main>

      {/* Footer */}
      <footer style={{ backgroundColor: "#f0f0f0", padding: "0.75rem", borderTop: "1px solid #ddd", textAlign: "center", fontSize: "0.9rem" }}>
        &copy; {new Date().getFullYear()} CST App
      </footer>
    </div>
  );
}
