package main

import (
	"fmt"
	"os"

	"github.com/GabrielSolioz/jenvx/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: jenvx <comando>")
		return
	}

	switch os.Args[1] {
	case "version":
		fmt.Println("jenvx 0.0.1")

	case "doctor":
		cmd.Doctor()

	default:
		fmt.Printf("Comando desconocido: %s\n", os.Args[1])
	}
}
