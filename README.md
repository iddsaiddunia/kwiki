# kwiki 🚀

Your cross-platform dev environment setup tool. One command to set up your entire dev environment on any machine.

## Quick Start (Bootstrap)

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/you/kwiki/main/scripts/bootstrap.ps1 | iex
```

**macOS/Linux:**
```bash
curl -fsSL https://raw.githubusercontent.com/you/kwiki/main/scripts/bootstrap.sh | bash
```

## Commands

| Command | Description |
|---------|-------------|
| `kwiki install` | Interactive checklist to select and install tools |
| `kwiki list` | List all available tools and versions |
| `kwiki export [file]` | Export tool list to `kwiki-env.yaml` |
| `kwiki import [file]` | Install tools from a saved `kwiki-env.yaml` |

## Available Tools

| Tool | Versions |
|------|----------|
| Python | 3.13, 3.12, 3.11, 3.10 |
| JDK | 21, 17, 11 |
| NVM | latest |
| Flutter | stable, beta, 3.27.0, 3.24.0 |
| Android Studio | latest, 2024.1, 2023.3 |
| PostgreSQL | 17, 16, 15, 14 |
| Git | latest |
| VS Code | latest |
| Docker | latest |
| Node.js | 22, 20, 18 |
| Go | 1.24, 1.23, 1.22 |
| Rust | stable, nightly |
| kubectl | latest |
| Terraform | latest, 1.10, 1.9 |
| Postman | latest |

## Save & Reuse Your Setup

```bash
# After setting up your machine, export your config
kwiki export my-setup.yaml

# On a new machine, just run
kwiki import my-setup.yaml
```

## Build from Source

```bash
git clone https://github.com/you/kwiki.git
cd kwiki
go build -o kwiki .
./kwiki install
```

## Cross-compile

```bash
# Windows binary from macOS/Linux
GOOS=windows GOARCH=amd64 go build -o kwiki.exe .

# macOS binary
GOOS=darwin GOARCH=amd64 go build -o kwiki-mac .

# Linux binary
GOOS=linux GOARCH=amd64 go build -o kwiki-linux .
```
