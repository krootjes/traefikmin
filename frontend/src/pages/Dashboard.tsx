import { useEffect, useState } from "react";

export default function Dashboard() {
  const [username, setUsername] = useState("");
  const [loading, setLoading] = useState(true);

  const handleLogout = () => {
    localStorage.removeItem("token");
    window.location.reload();
  };

  useEffect(() => {
    const fetchUser = async () => {
      const res = await fetch("/api/me", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });

      if (res.ok) {
        const data = await res.json();
        setUsername(data.username);
      } else {
        handleLogout();
      }

      setLoading(false);
    };

    fetchUser();
  }, []);

  if (loading) return <p style={styles.loading}>Bezig met laden...</p>;

  return (
    <div style={styles.container}>
      <h2>Welkom, {username} 👋</h2>
      <p>Je bent succesvol ingelogd met JWT-authenticatie.</p>
      <button onClick={handleLogout} style={styles.logout}>
        Uitloggen
      </button>
    </div>
  );
}

const styles: { [key: string]: React.CSSProperties } = {
  container: {
    maxWidth: 500,
    margin: "100px auto",
    padding: 24,
    border: "1px solid #ccc",
    borderRadius: 8,
    textAlign: "center",
    boxShadow: "0 2px 6px rgba(0,0,0,0.1)",
    fontFamily: "sans-serif",
  },
  logout: {
    marginTop: 20,
    padding: "10px 16px",
    backgroundColor: "#dc3545",
    color: "white",
    border: "none",
    borderRadius: 4,
    cursor: "pointer",
  },
  loading: {
    textAlign: "center",
    fontSize: "1.2rem",
    marginTop: "5rem",
  },
};
