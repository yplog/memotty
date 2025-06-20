#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

info() {
    echo -e "${BLUE}[INFO]${NC} $1" >&2
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1" >&2
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1" >&2
}

error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
    exit 1
}

detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)
    
    case $os in
        linux)
            OS="linux"
            ;;
        darwin)
            OS="darwin"
            ;;
        *)
            error "Unsupported operating system: $os"
            ;;
    esac
    
    case $arch in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            error "Unsupported architecture: $arch"
            ;;
    esac
    
    info "Detected platform: $OS-$ARCH"
}

get_latest_release() {
    info "Getting latest release information..."
    
    if command -v curl >/dev/null 2>&1; then
        LATEST_RELEASE=$(curl -s https://api.github.com/repos/yplog/memotty/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget >/dev/null 2>&1; then
        LATEST_RELEASE=$(wget -qO- https://api.github.com/repos/yplog/memotty/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        error "curl or wget is required"
    fi
    
    if [ -z "$LATEST_RELEASE" ]; then
        error "Could not fetch latest release information"
    fi
    
    info "Latest release: $LATEST_RELEASE"
}

download_binary() {
    local binary_name="memotty-${OS}-${ARCH}"
    local download_url="https://github.com/yplog/memotty/releases/download/${LATEST_RELEASE}/${binary_name}"
    local temp_file="/tmp/memotty"
    
    info "Downloading from: $download_url"
    
    if command -v curl >/dev/null 2>&1; then
        curl -L --progress-bar -o "$temp_file" "$download_url" || error "Download failed"
    elif command -v wget >/dev/null 2>&1; then
        wget --progress=bar:force -O "$temp_file" "$download_url" || error "Download failed"
    else
        error "curl or wget is required"
    fi
    
    if [ ! -f "$temp_file" ]; then
        error "Downloaded file not found"
    fi
    
    success "Download completed"
    
    echo "$temp_file"
}

install_binary() {
    local temp_file="$1"
    local install_dir="$HOME/.local/bin"
    local install_path="$install_dir/memotty"
    
    mkdir -p "$install_dir"
    
    cp "$temp_file" "$install_path"
    chmod +x "$install_path"
    
    rm -f "$temp_file"
    
    success "Binary installed to: $install_path"
    
    if [[ ":$PATH:" != *":$install_dir:"* ]]; then
        warning "⚠️  $install_dir is not in your PATH"
        info "Add this line to your ~/.bashrc or ~/.zshrc:"
        echo "" >&2
        echo "export PATH=\"\$HOME/.local/bin:\$PATH\"" >&2
        echo "" >&2
    fi
}

main() {    
    detect_platform
    get_latest_release
    
    local temp_file=$(download_binary)
    install_binary "$temp_file"
    
    echo "" >&2
    success "🎉 Installation completed successfully!"
    info "More info: https://github.com/yplog/memotty"
}

main "$@"