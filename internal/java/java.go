package java

import (
	"os"
	"os/exec"
	"strings"
)

type Info struct {
	PathVersion     string
	PathMajor       string
	JavaHome        string
	JavaHomeVersion string
	JavaHomeMajor   string
}

func Detect() Info {
	info := Info{
		JavaHome: os.Getenv("JAVA_HOME"),
	}

	info.PathVersion = getJavaVersion("java")
	info.PathMajor = ExtractMajor(info.PathVersion)

	if info.JavaHome != "" {
		javaExe := info.JavaHome + `\bin\java.exe`

		info.JavaHomeVersion = getJavaVersion(javaExe)
		info.JavaHomeMajor = ExtractMajor(info.JavaHomeVersion)
	}

	return info
}

func getJavaVersion(command string) string {
	cmd := exec.Command(command, "--version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")

	if len(lines) == 0 {
		return ""
	}

	return strings.TrimSpace(lines[0])
}

func ExtractMajor(versionLine string) string {
	fields := strings.Fields(versionLine)

	if len(fields) < 2 {
		return ""
	}

	version := fields[1]

	parts := strings.Split(version, ".")

	if len(parts) == 0 {
		return ""
	}

	return parts[0]
}
