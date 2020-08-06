package main

import (
	"github.com/jaxxstorm/http-blob-reader/cmd/http-blob-reader/serve"
	"github.com/jaxxstorm/http-blob-reader/cmd/http-blob-reader/version"
	"github.com/spf13/cobra"
	"os"
)

func configureCLI() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:  "http-role",
		Long: "A web server that displays the content of a blob store key",
	}

	rootCommand.AddCommand(serve.Command())
	rootCommand.AddCommand(version.Command())

	return rootCommand
}

func main() {
	rootCommand := configureCLI()

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
