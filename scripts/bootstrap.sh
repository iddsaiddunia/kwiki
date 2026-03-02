#!/usr/bin/env bash
# kwiki bootstrap for macOS/Linux
# Usage: curl -fsSL https://raw.githubusercontent.com/you/kwiki/main/scripts/bootstrap.sh | bash

set -e

echo "🚀 kwiki bootstrap starting..."

KWIKI_DIR="$HOME/.kwiki"

# Install Go if missing
if ! command -v go &>/dev/null; then
    echo "📦 Installing Go..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        if ! command -v brew &>/dev/null; then
            /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        fi
        brew install go
    else
        GO_VERSION="1.24.0"
        curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" | sudo tar -C /usr/local -xz
        export PATH=$PATH:/usr/local/go/bin
    fi
fi

# Install Git if missing
if ! command -v git &>/dev/null; then
    echo "📦 Installing Git..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        brew install git
    else
        sudo apt-get update -qq && sudo apt-get install -y git
    fi
fi

# Clone kwiki
if [ ! -d "$KWIKI_DIR" ]; then
    git clone https://github.com/you/kwiki.git "$KWIKI_DIR"
fi

# Build
cd "$KWIKI_DIR"
go build -o kwiki .

# Add to PATH
SHELL_RC="$HOME/.bashrc"
[[ "$SHELL" == *zsh* ]] && SHELL_RC="$HOME/.zshrc"

if ! grep -q "$KWIKI_DIR" "$SHELL_RC" 2>/dev/null; then
    echo "export PATH=\"\$PATH:$KWIKI_DIR\"" >> "$SHELL_RC"
fi
export PATH="$PATH:$KWIKI_DIR"

echo "✅ kwiki installed! Run: kwiki install"
kwiki install
