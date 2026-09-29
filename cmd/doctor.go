package cmd

import (
	"fmt"

	"github.com/GabrielSolioz/jenvx/internal/java"
	"github.com/GabrielSolioz/jenvx/internal/maven"
	"github.com/GabrielSolioz/jenvx/internal/project"
)

func Doctor() {
	fmt.Println("----------------")
	fmt.Println("- jenvx doctor -")
	fmt.Println("----------------")

	info := java.Detect()

	fmt.Println()
	fmt.Println("Java")

	if info.PathVersion == "" {
		fmt.Println("✗ Java was not found in PATH")
	} else {
		fmt.Println("✓ java detected")
		fmt.Println(" ", info.PathVersion)
	}

	fmt.Println()
	fmt.Println("JAVA_HOME")

	if info.JavaHome == "" {
		fmt.Println("⚠ JAVA_HOME is not defined")
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
		fmt.Println("⚠ Possible Java environment mismatch detected")
		fmt.Println("  Java in PATH:", info.PathVersion)
		fmt.Println("  JAVA_HOME:", info.JavaHomeVersion)
	}

	mavenInfo := maven.Detect()

	fmt.Println()
	fmt.Println("Maven")

	if !mavenInfo.Installed {
		fmt.Println("✗ Maven was not found in PATH")
	} else {
		fmt.Println("✓ Maven detected")
		fmt.Println(" ", mavenInfo.Version)
		fmt.Println(" ", mavenInfo.JavaVersion)
	}

	projectInfo := project.DetectPom()
	fmt.Println()
	fmt.Println("Project")

	if !projectInfo.IsMaven {
		fmt.Println("⚠  pom.xml was not found")
	} else {
		fmt.Println("✓ Maven Proyect detected")

		if projectInfo.JavaVersion != "" {
			fmt.Println("Required Java version:", projectInfo.JavaVersion)
		} else {
			fmt.Println("⚠ Could not determine the required Java version")
		}
	}
	fmt.Println()
	fmt.Println("Environment compatibility")

	if projectInfo.JavaVersion == "" {
		fmt.Println("⚠ Compatibility cannot be checked without a required Java version")
		return
	}

	required := projectInfo.JavaVersion

	fmt.Println("  Project requires Java", required)

	if info.JavaHomeMajor == required {
		fmt.Println("✓ JAVA_HOME matches project: Java", info.JavaHomeMajor)
	} else if info.JavaHomeMajor != "" {
		fmt.Printf("✗ JAVA_HOME uses Java %s, project requires Java %s\n",
			info.JavaHomeMajor,
			required,
		)
	}

	if mavenInfo.JavaMajor == required {
		fmt.Println("✓ Maven is using Java", mavenInfo.JavaMajor)
	} else if mavenInfo.JavaMajor != "" {
		fmt.Printf("✗ Maven is using Java %s, project requires Java %s\n",
			mavenInfo.JavaMajor,
			required,
		)
	}

	if info.PathMajor == required {
		fmt.Println("✓ PATH matches project: Java", info.PathMajor)
	} else if info.PathMajor != "" {
		fmt.Printf("⚠ PATH is using Java %s, project requires Java %s\n",
			info.PathMajor,
			required,
		)
	}
}
