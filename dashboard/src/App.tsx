import React, { useEffect, useState } from "react";
import axios from "axios";

interface Application {
  name: string;
  version: string;
}

interface CISCheck {
  name: string;
  passed: boolean;
  evidence: string;
}

interface HostData {
  hostname: string;
  applications: Application[];
  cis_checks: CISCheck[];
}

const App: React.FC = () => {
  const [data, setData] = useState<HostData | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchData = async () => {
    setLoading(true);
    setError(null);
    try {
      const res = await axios.get<HostData>("http://localhost:8080/host");
      setData(res.data);
    } catch (err) {
      setError("Failed to fetch data from agent");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  if (loading) return <div className="text-center mt-20 text-xl">⏳ Loading...</div>;
  if (error) return <div className="text-center mt-20 text-red-500">{error}</div>;
  if (!data) return <div className="text-center mt-20 text-xl">No data available.</div>;

  return (
    <div style={{ fontFamily: "Arial", padding: "20px" }}>
      <h1 style={{ textAlign: "center", color: "#2d3748" }}>🧩 Windows Agent Dashboard</h1>
      
      <div style={{ textAlign: "center", marginTop: "10px" }}>
        <button
          onClick={fetchData}
          style={{
            backgroundColor: "#3182ce",
            color: "#fff",
            padding: "8px 16px",
            borderRadius: "5px",
            cursor: "pointer",
            border: "none",
          }}
        >
          🔄 Sync Now
        </button>
      </div>

      <h2 style={{ marginTop: "20px" }}>
        Hostname: <span style={{ color: "#3182ce" }}>{data.hostname}</span>
      </h2>

      <ApplicationsTable applications={data.applications} />
      <CISTable cisChecks={data.cis_checks} />
    </div>
  );
};

// ---------------------------
// Reusable Components
// ---------------------------

const ApplicationsTable: React.FC<{ applications: Application[] }> = ({ applications }) => (
  <section style={{ marginTop: "30px" }}>
    <h3>Installed Applications</h3>
    <table border={1} cellPadding={6} style={{ width: "100%", marginTop: "10px" }}>
      <thead style={{ backgroundColor: "#edf2f7" }}>
        <tr>
          <th>Name</th>
          <th>Version</th>
        </tr>
      </thead>
      <tbody>
        {applications.map((app, i) => (
          <tr key={i}>
            <td>{app.name}</td>
            <td>{app.version}</td>
          </tr>
        ))}
      </tbody>
    </table>
  </section>
);

const CISTable: React.FC<{ cisChecks: CISCheck[] }> = ({ cisChecks }) => (
  <section style={{ marginTop: "40px" }}>
    <h3>CIS Security Checks</h3>
    <table border={1} cellPadding={6} style={{ width: "100%", marginTop: "10px" }}>
      <thead style={{ backgroundColor: "#edf2f7" }}>
        <tr>
          <th>Check</th>
          <th>Status</th>
          <th>Evidence</th>
        </tr>
      </thead>
      <tbody>
        {cisChecks.map((check, i) => (
          <tr key={i} style={{ backgroundColor: check.passed ? "#c6f6d5" : "#fed7d7" }}>
            <td>{check.name}</td>
            <td>{check.passed ? "✅ PASS" : "❌ FAIL"}</td>
            <td>{check.evidence}</td>
          </tr>
        ))}
      </tbody>
    </table>
  </section>
);

export default App;
