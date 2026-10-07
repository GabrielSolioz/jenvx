package main

import (
	"fmt"
	"os"

	"github.com/GabrielSolioz/jenvx/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jenvx <command>")
		return
	}

	switch os.Args[1] {
	case "version":
		fmt.Println("jenvx 0.0.1")

	case "doctor":
		cmd.Doctor()

	case "init":
		cmd.Init()

	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: jenvx run <command> [arguments...]")
			return
		}

		cmd.Run(os.Args[2], os.Args[3:])

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
