package cmd

import (
	"fmt"
	"os"

	"github.com/GabrielSolioz/jenvx/internal/project"
)

func Init() {
	if _, err := os.Stat("jenvx.toml"); err == nil {
		fmt.Println("✗ jenvx.toml already exists")
		return
	}

	projectInfo := project.DetectPom()

	javaVersion := "21"
	buildTool := "maven"

	if projectInfo.IsMaven && projectInfo.JavaVersion != "" {
		javaVersion = projectInfo.JavaVersion
	}

	content := fmt.Sprintf(`[java]
version = "%s"

[build]
tool = "%s"

[commands]
dev = "./mvnw spring-boot:run"
test = "./mvnw test"
build = "./mvnw clean package"
`, javaVersion, buildTool)

	err := os.WriteFile("jenvx.toml", []byte(content), 0644)
	if err != nil {
		fmt.Println("✗ Failed to create jenvx.toml:", err)
		return
	}

	fmt.Println("✓ Created jenvx.toml")
	fmt.Println("  Java:", javaVersion)
	fmt.Println("  Build tool:", buildTool)
}
