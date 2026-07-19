# Mini POS — Desktop Application

A lightning-fast, offline-first Desktop Point of Sale (POS) application built with **Go**, **Wails v2**, **Vue 3**, and **SQLite**. 

Because this project utilizes a **pure-Go SQLite driver** (`glebarez/sqlite`), it does not require CGO. This makes the application incredibly easy to cross-compile across operating systems!

---

## 🛠️ Prerequisites

Before you begin, ensure you have the following installed on your machine:

1. **Go (1.24+)**: [Download Go](https://go.dev/dl/)
2. **Node.js & NPM**: [Download Node.js](https://nodejs.org/) (v16+ recommended)
3. **Wails CLI**: Install the Wails v2 build tool by running:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

### 🐧 Linux Specific Dependencies
If you are developing or building on Linux (Ubuntu/Debian), Wails relies on native WebKit libraries to render the UI window. You **must** install these packages:

```bash
# Update package list and install WebKitGTK headers
sudo apt-get update --allow-releaseinfo-change
sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.0-dev build-essential
```

*(Note: If you run into release-info-change errors during `apt-get update`, ensure you use the `--allow-releaseinfo-change` flag as written above).*

---

## 💻 Local Development

To run the application in live-reload mode (where changes to Vue or Go files update instantly without having to recompile the app manually):

1. Open your terminal in the project root (`/mini-pos`).
2. Run the Wails development server:
   ```bash
   wails dev
   ```
This command will compile the Go backend, start the Vite frontend server, and open the native desktop window.

> **Frontend UI Only**: If you only want to tweak the Vue UI in a browser without running the Go backend, you can navigate to the `frontend/` folder and run `npm run dev`.

---

## 📦 Building for Production

When you are ready to create a standalone, production-ready executable for your users:

### 1. Build for your Current OS (Linux)
Run this from the project root:
```bash
wails build
```
Once completed, you will find your standalone application binary in the `build/bin/` folder.

### 2. Cross-Compile for Windows (from Linux)
Because we bypassed CGO by using the pure-Go SQLite driver, cross-compiling to Windows from Linux is incredibly simple.

First, ensure you have the Windows cross-compiler tools installed on your Linux machine:
```bash
sudo apt install -y mingw-w64
```

Then, instruct Wails to build for Windows:
```bash
wails build -platform windows/amd64
```
Your final `pos-mini.exe` file will be generated and placed inside the `build/bin/` folder, ready to be copied and run natively on any Windows computer!

---

## 🗄️ Database & Storage
- The application automatically generates an offline SQLite database inside the `data/pos.db` folder on startup.
- Application logs are written securely to `logs/app.log`.
- **Backups**: You can trigger safe database backups directly from the App's **Settings** interface.
