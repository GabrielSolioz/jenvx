package java

import (
	"os"
	"os/exec"
	"strings"
)

type Info struct {
	PathVersion     string
	JavaHome        string
	JavaHomeVersion string
}

func Detect() Info {
	info := Info{
		JavaHome: os.Getenv("JAVA_HOME"),
	}

	info.PathVersion = getJavaVersion("java")

	if info.JavaHome != "" {
		javaExe := info.JavaHome + `\bin\java.exe`
		info.JavaHomeVersion = getJavaVersion(javaExe)
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
