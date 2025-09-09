// internal/cli/cli_test.go
package cli

import (
	"testing"
)

func TestNewRootCommand(t *testing.T) {
	// Create the root command
	rootCmd := NewRootCommand()
	
	// Check that the root command has the expected subcommands
	subcommands := []string{
		"validate",
		"create",
		"run",
		"deploy",
		"mermaid",
		"meta-agents",
		"serve",
		"clean",
		"create-cr",
	}
	
	for _, name := range subcommands {
		cmd, _, err := rootCmd.Find([]string{name})
		if err != nil {
			t.Errorf("Could not find subcommand %s: %v", name, err)
		}
		
		if cmd.Name() != name {
			t.Errorf("Expected command name %s, got %s", name, cmd.Name())
		}
	}
}

