package maven

import (
	"os"
	"os/exec"
	"strings"
)

type Info struct {
	Installed       bool
	Version         string
	JavaVersion     string
	JavaMajor       string
	ProjectDetected bool
}

func Detect() Info {
	info := Info{}

	info.Version, info.JavaVersion = getMavenInfo()
	info.JavaMajor = extractJavaMajor(info.JavaVersion)

	if info.Version != "" {
		info.Installed = true
	}

	if _, err := os.Stat("pom.xml"); err == nil {
		info.ProjectDetected = true
	}

	return info
}

func getMavenInfo() (string, string) {
	cmd := exec.Command("mvn", "--version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", ""
	}

	lines := strings.Split(string(output), "\n")

	var version string
	var javaVersion string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Apache Maven") {
			version = line
		}

		if strings.HasPrefix(line, "Java version:") {
			javaVersion = line
		}
	}

	return version, javaVersion
}

func extractJavaMajor(line string) string {
	line = strings.TrimPrefix(line, "Java version:")
	line = strings.TrimSpace(line)

	parts := strings.Split(line, ",")

	if len(parts) == 0 {
		return ""
	}

	version := strings.TrimSpace(parts[0])

	versionParts := strings.Split(version, ".")

	if len(versionParts) == 0 {
		return ""
	}

	return versionParts[0]
}
