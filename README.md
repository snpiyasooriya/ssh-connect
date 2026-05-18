# 🖥️ SSH Connect

An elegant, high-performance desktop application for managing and launching SSH connections. Built with **Go**, **Wails v2**, and modern **Vanilla JS/CSS3**, **SSH Connect** reads your local `~/.ssh/config` file, presents a beautifully designed responsive grid of your hosts, and lets you connect instantly by spawning active terminal sessions in **Alacritty**.

> [!NOTE]
> This application leverages **Wails v2** to bundle a lightweight Go backend with an ultra-responsive web frontend, delivering a seamless native-desktop experience.

---

## ✨ Features

- **📂 Native SSH Configuration Sync**: Automatically reads and parses your existing `~/.ssh/config` file, mapping all host configurations instantly without manual setup.
- **🔍 Instant Search & Filter**: Interactive real-time fuzzy filtering lets you search through dozens of hostnames on-the-fly.
- **⚡ One-Click Terminal Launch**: Selecting any host card instantly spawns a hardware-accelerated **Alacritty** terminal emulator running `ssh <host>`.
- **🎨 Premium Dark UI**: Designed with a high-fidelity visual experience featuring modern typography, rich gradient headings, elegant glassmorphism backgrounds (`backdrop-filter`), smooth hover states, and dynamic entry-fade micro-animations.

---

## 🛠️ Architecture & Tech Stack

### Backend (Go)
- **Framework**: [Wails v2](https://wails.io) for building lightweight desktop apps with Go.
- **Core Logic ([app.go](file:///home/sayuru/ratko/ssh-connect/app.go))**:
  - `GetSSHHosts()`: Scans and parses `~/.ssh/config` line-by-line, dynamically capturing custom aliases and filtering out wildcards (`*`) and empty declarations.
  - `ConnectToHost(host)`: Safely invokes native terminal commands (`alacritty -e ssh <host>`) in a detached execution flow.

### Frontend (Web Integration)
- **Tech Stack**: Vanilla HTML5, JavaScript (ES6+), and CSS3 Custom Properties.
- **Design System ([style.css](file:///home/sayuru/ratko/ssh-connect/frontend/src/style.css))**: Crafted from scratch using modern design principles:
  - Curated, dark-mode color scheme with glowing cyan accents (`#00d2ff`).
  - Seamless CSS variables for cohesive style tokens.
  - Custom scrollbars, layout grid configurations, and responsive padding.
- **Interactive UI ([main.js](file:///home/sayuru/ratko/ssh-connect/frontend/src/main.js))**: Asynchronously requests SSH host details from Go during startup and binds responsive click-handlers that bridge frontend user actions with Go application runtime commands.

---

## 🚀 Getting Started

### 📋 Prerequisites

Ensure you have the following installed on your local machine:
1. **Go**: Version `1.23` or higher.
2. **Node.js & npm**: For managing and building frontend web assets.
3. **Wails CLI**: Install globally using:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```
4. **Alacritty**: Installed and available in your environment's system `PATH` (used as the default terminal execution environment).
5. **SSH Config**: Ensure you have some valid hosts defined in your local configuration at `~/.ssh/config`. For example:
   ```ssh
   Host production-server
       HostName 198.51.100.1
       User ubuntu
       IdentityFile ~/.ssh/id_rsa

   Host staging-db web-node
       HostName 198.51.100.2
       User deploy
   ```

---

## 💻 Development & Deployment

### Run in Development Mode (Live Reload)

To start the application in development mode with active hot-reloading for both backend Go code and frontend web code:

```bash
wails dev
```

> [!TIP]
> While running in development mode, a separate Vite server will automatically start. You can also debug and test your interface from any web browser by navigating to: **`http://localhost:34115`**.

### Building for Production

To compile a highly optimized, single binary executable for your platform, run:

```bash
wails build
```

The resulting executable binary will be generated under the `build/bin/` directory.

---

## 📁 Repository Structure

```txt
ssh-connect/
├── app.go             # Backend application logic (SSH config parser & command runner)
├── app_test.go        # Unit testing suite for Go parsing functions
├── main.go            # Entry point of the Wails app (window & setup config)
├── wails.json         # Wails project configuration settings
├── go.mod             # Go module definition & dependency map
├── build/             # Output folder for production binaries & assets
└── frontend/          # Web frontend source code
    ├── index.html     # Single Page Application HTML root template
    ├── package.json   # Frontend dependency manifest
    └── src/
        ├── main.js    # Logic handling search, data binding, and bridge calls
        ├── style.css  # Premium layout styles, custom scrollbars, and keyframes
        └── assets/    # Frontend image assets and icon directories
```

---

## 🧪 Testing

To run unit tests validating the Go backend parsing logic:

```bash
go test -v ./...
```

---

## 📄 License & Creator

- **Author**: Sayuru Piyasooriya ([snpiyasooriya@gmail.com](mailto:snpiyasooriya@gmail.com))
