import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";

export default function App() {
  const token = localStorage.getItem("token");

  return (
    <div style={styles.app}>
      <h1 style={styles.title}>Traefikmin</h1>
      {token ? <Dashboard /> : <Login />}
    </div>
  );
}

const styles: { [key: string]: React.CSSProperties } = {
  app: {
    fontFamily: "sans-serif",
    padding: "2rem",
    textAlign: "center",
  },
  title: {
    marginBottom: "2rem",
  },
};
