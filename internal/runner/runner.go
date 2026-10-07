package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/GabrielSolioz/jenvx/internal/config"
	"github.com/GabrielSolioz/jenvx/internal/java"
)

func Run(command string, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("could not load jenvx.toml: %w", err)
	}

	javaInfo := java.Detect()

	if javaInfo.JavaHome == "" {
		return fmt.Errorf("JAVA_HOME is not defined")
	}

	if javaInfo.JavaHomeMajor == "" {
		return fmt.Errorf("could not determine JAVA_HOME Java version")
	}

	requiredJava := cfg.Java.Version

	if requiredJava == "" {
		return fmt.Errorf("Java version is not defined in jenvx.toml")
	}

	if javaInfo.JavaHomeMajor != requiredJava {
		return fmt.Errorf(
			"JAVA_HOME uses Java %s, but project requires Java %s",
			javaInfo.JavaHomeMajor,
			requiredJava,
		)
	}

	javaBin := filepath.Join(javaInfo.JavaHome, "bin")

	env := os.Environ()

	env = setEnv(env, "JAVA_HOME", javaInfo.JavaHome)
	env = prependPath(env, javaBin)

	resolvedCommand := resolveCommand(command, javaBin)

	execCmd := exec.Command(resolvedCommand, args...)

	execCmd.Env = env
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	fmt.Printf("jenvx: using Java %s\n", requiredJava)
	fmt.Printf("jenvx: JAVA_HOME=%s\n", javaInfo.JavaHome)
	fmt.Println()

	return execCmd.Run()
}

func setEnv(env []string, key string, value string) []string {
	entry := key + "=" + value

	for i, current := range env {
		parts := strings.SplitN(current, "=", 2)

		if len(parts) != 2 {
			continue
		}

		if strings.EqualFold(parts[0], key) {
			env[i] = entry
			return env
		}
	}

	return append(env, entry)
}

func prependPath(env []string, directory string) []string {
	for i, current := range env {
		parts := strings.SplitN(current, "=", 2)

		if len(parts) != 2 {
			continue
		}

		if strings.EqualFold(parts[0], "PATH") {
			env[i] = "PATH=" + directory + string(os.PathListSeparator) + parts[1]
			return env
		}
	}

	return append(env, "PATH="+directory)
}

func resolveCommand(command string, javaBin string) string {
	// If the user provided a full or relative path, don't modify it.
	if strings.ContainsAny(command, `/\`) {
		return command
	}

	candidates := []string{
		filepath.Join(javaBin, command),
		filepath.Join(javaBin, command+".exe"),
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	return command
}
