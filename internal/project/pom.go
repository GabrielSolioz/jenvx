package project

import (
	"encoding/xml"
	"os"
	"strings"
)

type Pom struct {
	XMLName    xml.Name   `xml:"project"`
	Properties Properties `xml:"properties"`
}

type Properties struct {
	JavaVersion     string `xml:"java.version"`
	CompilerRelease string `xml:"maven.compiler.release"`
	CompilerSource  string `xml:"maven.compiler.source"`
}

type Info struct {
	IsMaven     bool
	JavaVersion string
}

func DetectPom() Info {
	data, err := os.ReadFile("pom.xml")
	if err != nil {
		return Info{}
	}

	var pom Pom

	if err := xml.Unmarshal(data, &pom); err != nil {
		return Info{
			IsMaven: true,
		}
	}

	javaVersion := detectJavaVersion(pom.Properties)

	return Info{
		IsMaven:     true,
		JavaVersion: javaVersion,
	}
}

func detectJavaVersion(properties Properties) string {
	if strings.TrimSpace(properties.JavaVersion) != "" {
		return strings.TrimSpace(properties.JavaVersion)
	}

	if strings.TrimSpace(properties.CompilerRelease) != "" {
		return strings.TrimSpace(properties.CompilerRelease)
	}

	if strings.TrimSpace(properties.CompilerSource) != "" {
		return strings.TrimSpace(properties.CompilerSource)
	}

	return ""
}
