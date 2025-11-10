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

  useEffect(() => {
    axios.get<HostData>("http://localhost:8080/host")
      .then(res => setData(res.data))
      .catch(err => console.error(err));
  }, []);

  if (!data) return <div className="text-center mt-20 text-xl">Loading...</div>;

  return (
    <div style={{ fontFamily: "Arial", padding: "20px" }}>
      <h1 style={{ textAlign: "center", color: "#2d3748" }}>🧩 Windows Agent Dashboard</h1>
      <h2 style={{ marginTop: "10px" }}>Hostname: <span style={{ color: "#3182ce" }}>{data.hostname}</span></h2>

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
            {data.applications.map((app, i) => (
              <tr key={i}>
                <td>{app.name}</td>
                <td>{app.version}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>

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
            {data.cis_checks.map((check, i) => (
              <tr key={i} style={{ backgroundColor: check.passed ? "#c6f6d5" : "#fed7d7" }}>
                <td>{check.name}</td>
                <td>{check.passed ? "✅ PASS" : "❌ FAIL"}</td>
                <td>{check.evidence}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>
    </div>
  );
};

export default App;
