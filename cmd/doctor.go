package cmd

import (
	"fmt"

	"github.com/GabrielSolioz/jenvx/internal/java"
)

func Doctor() {
	fmt.Println("jenvx doctor")
	fmt.Println("-------------")

	info := java.Detect()

	fmt.Println()
	fmt.Println("Java")

	if info.PathVersion == "" {
		fmt.Println("✗ Java no encontrado en PATH")
	} else {
		fmt.Println("✓ java encontrado")
		fmt.Println(" ", info.PathVersion)
	}

	fmt.Println()
	fmt.Println("JAVA_HOME")

	if info.JavaHome == "" {
		fmt.Println("⚠ JAVA_HOME no está definido")
	} else {
		fmt.Println("✓", info.JavaHome)

		if info.JavaHomeVersion != "" {
			fmt.Println(" ", info.JavaHomeVersion)
		}
	}

	if info.PathVersion != "" &&
		info.JavaHomeVersion != "" &&
		info.PathVersion != info.JavaHomeVersion {

		fmt.Println()
		fmt.Println("⚠ Posible inconsistencia detectada")
		fmt.Println("  Java en PATH:", info.PathVersion)
		fmt.Println("  JAVA_HOME:", info.JavaHomeVersion)
	}
}
