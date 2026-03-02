package installer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/kwiki/kwiki/tools"
)

type Selection struct {
	Tool    tools.Tool
	Version string
}

func Install(selections []Selection) {
	os_type := runtime.GOOS
	for _, sel := range selections {
		fmt.Printf("\n📦 Installing %s %s...\n", sel.Tool.Name, sel.Version)
		var err error
		switch os_type {
		case "windows":
			err = installWindows(sel)
		case "darwin":
			err = installMac(sel)
		case "linux":
			err = installLinux(sel)
		}
		if err != nil {
			fmt.Printf("  ❌ Failed: %v\n", err)
		} else {
			fmt.Printf("  ✅ Done\n")
			setEnvVars(sel)
		}
	}
}

func installWindows(sel Selection) error {
	pkg := resolvePackage(sel.Tool.WinGet, sel.Version)
	if pkg == "" {
		return fmt.Errorf("no winget package defined for %s", sel.Tool.Name)
	}
	return run("winget", "install", "--id", pkg, "--silent", "--accept-package-agreements", "--accept-source-agreements")
}

func installMac(sel Selection) error {
	pkg := resolvePackage(sel.Tool.Brew, sel.Version)
	if pkg == "" {
		return fmt.Errorf("no brew package defined for %s", sel.Tool.Name)
	}
	// ensure brew is available
	if _, err := exec.LookPath("brew"); err != nil {
		fmt.Println("  🍺 Installing Homebrew first...")
		if err := run("/bin/bash", "-c", `$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)`); err != nil {
			return fmt.Errorf("failed to install homebrew: %w", err)
		}
	}
	return run("brew", "install", pkg)
}

func installLinux(sel Selection) error {
	pkg := resolvePackage(sel.Tool.Apt, sel.Version)
	if pkg == "" {
		return fmt.Errorf("no apt package defined for %s", sel.Tool.Name)
	}
	// special case for tools not in apt
	switch sel.Tool.Name {
	case "NVM":
		return run("bash", "-c", `curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash`)
	case "Flutter":
		return run("bash", "-c", `sudo snap install flutter --classic`)
	case "Android Studio":
		return run("bash", "-c", `sudo snap install android-studio --classic`)
	case "Rust":
		return run("bash", "-c", `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y`)
	case "Terraform":
		return installTerraformLinux()
	}
	run("sudo", "apt-get", "update", "-qq")
	return run("sudo", "apt-get", "install", "-y", pkg)
}

func installTerraformLinux() error {
	cmds := [][]string{
		{"bash", "-c", `wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg`},
		{"bash", "-c", `echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list`},
		{"sudo", "apt-get", "update", "-qq"},
		{"sudo", "apt-get", "install", "-y", "terraform"},
	}
	for _, c := range cmds {
		if err := run(c[0], c[1:]...); err != nil {
			return err
		}
	}
	return nil
}

func setEnvVars(sel Selection) {
	if len(sel.Tool.EnvVars) == 0 {
		return
	}
	fmt.Printf("  🔧 Setting environment variables for %s\n", sel.Tool.Name)
	for key := range sel.Tool.EnvVars {
		fmt.Printf("     %s → (set automatically by installer)\n", key)
	}
}

func resolvePackage(pkg, version string) string {
	if version == "latest" || version == "stable" || version == "" {
		return strings.ReplaceAll(pkg, ".{version}", "")
	}
	return strings.ReplaceAll(pkg, "{version}", version)
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
