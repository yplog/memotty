package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/yplog/memotty/internal/ui"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestVersion() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/yplog/memotty/releases/latest")
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest release: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var release GitHubRelease
	if err := json.Unmarshal(body, &release); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	return release.TagName, nil
}

func normalizeVersion(version string) string {
	return strings.TrimPrefix(version, "v")
}

func Run() error {
	fmt.Println("🔍 Checking for updates...")

	currentVersion := ui.GetVersionInfo()
	currentVersionClean := normalizeVersion(strings.TrimPrefix(currentVersion, "memotty "))

	if currentVersionClean == "dev" {
		fmt.Println("Development version detected, proceeding with update...")
		return runInstallScript()
	}

	latestVersion, err := getLatestVersion()
	if err != nil {
		fmt.Printf("Could not check latest version: %v\n", err)
		fmt.Println("Proceeding with update anyway...")
		return runInstallScript()
	}

	latestVersionClean := normalizeVersion(latestVersion)

	fmt.Printf("Current version: %s\n", currentVersionClean)
	fmt.Printf("Latest version: %s\n", latestVersionClean)

	if currentVersionClean >= latestVersionClean {
		fmt.Println("Already up to date! No update needed.")
		return nil
	}

	fmt.Printf("Updating from %s to %s...\n", currentVersionClean, latestVersionClean)
	return runInstallScript()
}

func runInstallScript() error {
	cmd := exec.Command("bash", "-c", "curl -fsSL https://raw.githubusercontent.com/yplog/memotty/main/scripts/install.sh | bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
