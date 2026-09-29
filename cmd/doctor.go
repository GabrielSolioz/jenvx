package cmd

import (
	"fmt"

	"github.com/GabrielSolioz/jenvx/internal/java"
	"github.com/GabrielSolioz/jenvx/internal/maven"
	"github.com/GabrielSolioz/jenvx/internal/project"
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

	mavenInfo := maven.Detect()

	fmt.Println()
	fmt.Println("Maven")

	if !mavenInfo.Installed {
		fmt.Println("✗ Maven no está instalado o no está en PATH")
	} else {
		fmt.Println("✓ Maven detectado")
		fmt.Println(" ", mavenInfo.Version)
		fmt.Println(" ", mavenInfo.JavaVersion)
	}

	projectInfo := project.DetectPom()
	fmt.Println()
	fmt.Println("Project")

	if !projectInfo.IsMaven {
		fmt.Println("⚠ No se encontró pom.xml")
	} else {
		fmt.Println("✓ Proyecto Maven detectado")

		if projectInfo.JavaVersion != "" {
			fmt.Println("  Java requerido:", projectInfo.JavaVersion)
		} else {
			fmt.Println("⚠ No se pudo determinar la versión de Java requerida")
		}
	}
}
