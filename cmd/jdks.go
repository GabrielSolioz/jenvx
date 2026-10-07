package cmd

import (
	"fmt"

	"github.com/GabrielSolioz/jenvx/internal/jdk"
)

func JDKs() {
	baseDir, err := jdk.BaseDir()
	if err != nil {
		fmt.Println("✗ Failed to determine jenvx JDK directory:", err)
		return
	}

	fmt.Println("jenvx JDK store")
	fmt.Println("  Location:", baseDir)

	if err := jdk.EnsureBaseDir(); err != nil {
		fmt.Println("✗ Failed to create JDK store:", err)
		return
	}

	fmt.Println("✓ JDK store ready")
}
