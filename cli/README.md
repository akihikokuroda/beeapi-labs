# Maestro CLI

The Maestro CLI is a command-line tool for managing Maestro workflows. It provides commands for creating, running, and deploying workflows.

## Installation

To install the Maestro CLI, you need to have Go 1.24 or later installed. Then you can install the CLI using the following command:

```bash
go install github.com/ai4quantum/maestro/cli/cmd/maestro@latest


Usage
The Maestro CLI provides the following commands:

validate: Validate YAML files against JSON schemas
create: Create agents from a configuration file
run: Run a workflow with specified agents and workflow files
deploy: Deploy a workflow to a Kubernetes cluster or local server
mermaid: Generate mermaid diagrams from a workflow file
meta-agents: Run meta-agents on a text file
serve: Serve agents via HTTP endpoints
clean: Clean up running processes
create-cr: Create Kubernetes custom resources
For more information on each command, use the --help flag:

maestro --help
maestro validate --help
maestro create --help
# etc.

bash


Examples
Validate a YAML file
maestro validate schema.json file.yaml

bash


Create agents from a configuration file
maestro create agents.yaml

bash


Run a workflow
maestro run agents.yaml workflow.yaml

bash


Deploy a workflow to Kubernetes
maestro deploy agents.yaml workflow.yaml --k8s

bash


Generate a mermaid diagram
maestro mermaid workflow.yaml

bash


Serve an agent via HTTP
maestro serve agents.yaml --agent-name my-agent

bash


Clean up running processes
maestro clean

bash



## Migration Guide

Finally, let's create a migration guide to help users migrate from the Python CLI to the Go CLI:

```markdown
# Migrating from Python CLI to Go CLI

This guide will help you migrate from the Python CLI to the Go CLI.

## Installation

### Python CLI

```bash
pip install maestro-cli

txt


Go CLI
go install github.com/ai4quantum/maestro/cli/cmd/maestro@latest

bash


Command Mapping
The Go CLI provides the same commands as the Python CLI, with the same syntax and options. Here's a mapping of the commands:

Python CLI	Go CLI
maestro validate	maestro validate
maestro create	maestro create
maestro run	maestro run
maestro deploy	maestro deploy
maestro mermaid	maestro mermaid
maestro meta-agents	maestro meta-agents
maestro serve	maestro serve
maestro clean	maestro clean
maestro create-cr	maestro create-cr
Differences
While the Go CLI aims to be a drop-in replacement for the Python CLI, there are some differences to be aware of:

Performance: The Go CLI is generally faster than the Python CLI, especially for operations that involve file I/O or process management.

Dependencies: The Go CLI is a single binary with no external dependencies, while the Python CLI requires Python and several Python packages.

Error Handling: The Go CLI may provide more detailed error messages in some cases.

Output Formatting: The output formatting may be slightly different between the two CLIs, but the content should be the same.

Migration Steps
Install the Go CLI as described above.

Test your existing workflows with the Go CLI to ensure they work as expected.

Update any scripts or automation that use the Python CLI to use the Go CLI instead.

If you encounter any issues, please report them on the GitHub repository.


