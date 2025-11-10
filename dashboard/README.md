🧩 Windows Agent Dashboard

A lightweight Windows Security & Compliance Dashboard that visualizes host data collected by a Go-based agent.
The system performs CIS Windows 10/11 Level 1 Benchmark checks, collects installed applications, and exposes them through a REST API for visualization in a React dashboard.

🚀 Features
🖥️ Go Agent (Backend)

Collects Hostname

Enumerates Installed Applications

Performs 10+ CIS Security Checks, such as:

Firewall profiles enabled

BitLocker status

SMBv1 disabled

RDP NLA enabled

Password & Account Lockout policy

UAC enabled

Defender active

LSA Protection

Audit Policy Checks

Secure Boot & Windows Updates

Exposes results via REST endpoint:

GET http://localhost:8080/host

💻 React Dashboard (Frontend)

Displays:

Hostname

Installed Applications

CIS Check Results (✅ PASS / ❌ FAIL with evidence)

Provides “Sync Now” button to refresh data from the backend.

Uses Axios for API communication.

Responsive, minimal design using pure CSS.

🏗️ Architecture Overview
[ Windows Host ]
       |
       |--- Go Agent
       |        └── Collects system info & CIS checks
       |        └── Serves JSON at /host
       |
       └── React Dashboard (http://localhost:3000)
                └── Fetches and visualizes the data

⚙️ Installation & Setup
1️⃣ Backend (Go Agent)
git clone https://github.com/shaswatprakash/windows-agent.git
cd windows-agent

# Install dependencies
go mod tidy

# Run locally
go run main.go

# The agent will start on:
# http://localhost:8080/host

2️⃣ Frontend (React Dashboard)
git clone https://github.com/shaswatprakash/windows-dashboard.git
cd windows-dashboard

# Install dependencies
npm install

# Run in development mode
npm start


📍 Open http://localhost:3000
 in your browser.

The dashboard will automatically fetch host data from the Go agent.

🧪 API Example

Endpoint:

GET http://localhost:8080/host


Sample Response:

{
  "hostname": "DESKTOP-12345",
  "applications": [
    { "name": "Google Chrome", "version": "127.0.6533.78" },
    { "name": "Visual Studio Code", "version": "1.92.0" }
  ],
  "cis_checks": [
    { "name": "Firewall Enabled", "passed": true, "evidence": "All profiles active" },
    { "name": "SMBv1 Disabled", "passed": false, "evidence": "Feature detected" }
  ]
}

## Project Structure
windows-agent/
├── main.go                  # Entry point for Go agent
├── internal/
│   ├── collector/           # CIS and app inventory logic
│   └── sender/              # (Optional) network sender stub
└── go.mod / go.sum

windows-dashboard/
├── src/
│   ├── App.tsx              # Main React component
│   ├── components/          # Tables & UI components
│   └── index.tsx
└── package.json

### Building the Agent for Windows (from macOS)

You can cross-compile your Go code to generate a Windows executable directly on macOS.

Step 1️⃣ — Run this command from the project root
GOOS=windows GOARCH=amd64 go build -o windows-agent.exe main.go


💡 This creates a windows-agent.exe file inside your project directory.

Step 2️⃣ — Transfer to Windows

Copy windows-agent.exe to your Windows machine via USB or network share, then run:

.\windows-agent.exe


The agent will start and expose:

http://localhost:8080/host
👨‍💻 Author

Shaswat Prakash
Senior Software Developer / Full Stack Engineer
🌐 GitHub