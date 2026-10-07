package cmd

import (
	"fmt"

	"github.com/GabrielSolioz/jenvx/internal/runner"
)

func Run(command string, args []string) {
	err := runner.Run(command, args)

	if err != nil {
		fmt.Println("✗ Failed to run command:", err)
	}
}
