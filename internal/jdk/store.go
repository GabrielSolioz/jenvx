package jdk

import (
	"fmt"
	"os"
	"path/filepath"
)

func BaseDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine user home directory: %w", err)
	}

	return filepath.Join(home, ".jenvx", "jdks"), nil
}

func Path(version string) (string, error) {
	baseDir, err := BaseDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, version), nil
}

func Exists(version string) bool {
	jdkPath, err := Path(version)
	if err != nil {
		return false
	}

	javaExecutable := filepath.Join(jdkPath, "bin", javaExecutableName())

	if _, err := os.Stat(javaExecutable); err != nil {
		return false
	}

	return true
}

func EnsureBaseDir() error {
	baseDir, err := BaseDir()
	if err != nil {
		return err
	}

	return os.MkdirAll(baseDir, 0755)
}

func javaExecutableName() string {
	if os.PathSeparator == '\\' {
		return "java.exe"
	}

	return "java"
}
