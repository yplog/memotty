package ui

import "fmt"

var (
	appVersion   = "dev"
	appCommit    = "unknown"
	appBuildTime = "unknown"
)

func SetVersionInfo(version, commit, buildTime string) {
	appVersion = version
	appCommit = commit
	appBuildTime = buildTime
}

func GetVersionInfo() string {
	if appVersion == "dev" {
		return "memotty dev"
	}
	return fmt.Sprintf("memotty %s", appVersion)
}

func GetDetailedVersionInfo() string {
	return fmt.Sprintf("%s (commit: %s, built: %s)", appVersion, appCommit, appBuildTime)
}
