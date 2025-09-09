// internal/cli/cli.go
package cli

import (
	"github.com/spf13/cobra"
	
	"github.com/ai4quantum/maestro/cli/internal/commands"
)

// Global flags
var (
	verbose   bool
	silent    bool
	dryRun    bool
	mcpServerURI string
)

// NewRootCommand creates the root command for the Maestro CLI
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "maestro",
		Short:   "Maestro CLI for workflow management",
		Version: "0.0.4", // Same version as in the Python CLI
		Long: `Maestro CLI for workflow management.
		
This CLI provides commands for creating, running, and managing Maestro workflows.`,
	}

	// Add global flags
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Show all output")
	rootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "Show no additional output on success")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Mocks agents and other parts of workflow execution")
	rootCmd.PersistentFlags().StringVar(&mcpServerURI, "mcp-server-uri", "", "MCP server URI (overrides MAESTRO_KNOWLEDGE_MCP_SERVER_URI environment variable)")

	// Add commands
	rootCmd.AddCommand(
		commands.NewValidateCommand(),
		commands.NewCreateCommand(),
		commands.NewRunCommand(),
		commands.NewDeployCommand(),
		commands.NewMermaidCommand(),
		commands.NewMetaAgentsCommand(),
		commands.NewServeCommand(),
		commands.NewCleanCommand(),
		commands.NewCreateCrCommand(),
	)

	return rootCmd
}

