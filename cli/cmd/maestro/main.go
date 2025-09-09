// cmd/maestro/main.go
package main

import (
	"fmt"
	"os"

	"github.com/ai4quantum/maestro/cli/internal/cli"
	"github.com/ai4quantum/maestro/cli/internal/common"
)

func main() {
	// Load environment variables from .env file
	common.LoadEnv()

	// Create and execute the root command
	rootCmd := cli.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

