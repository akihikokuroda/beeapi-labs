// internal/cli/router.go
package cli

import (
	"fmt"
	
	"github.com/spf13/cobra"
	
	"github.com/ai4quantum/maestro/cli/internal/commands"
)

// ExecuteCommand executes a command and returns the result
func ExecuteCommand(cmd *cobra.Command, args []string) error {
	// Get the command name
	cmdName := cmd.Name()
	
	// Create command options
	options := commands.NewCommandOptions(cmd)
	
	// Route to the appropriate command handler
	var err error
	switch cmdName {
	case "validate":
		validateCmd := &commands.ValidateCommand{
			BaseCommand: commands.NewBaseCommand(options),
		}
		err = validateCmd.Run()
	case "create":
		createCmd := &commands.CreateCommand{
			BaseCommand: commands.NewBaseCommand(options),
		}
		err = createCmd.Run()
	// Add other commands here...
	default:
		err = fmt.Errorf("unknown command: %s", cmdName)
	}
	
	return err
}

