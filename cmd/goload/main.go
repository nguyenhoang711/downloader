package main

import (
	"fmt"
	"log"

	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/wiring"
	"github.com/spf13/cobra"
)

var (
	version    string
	commitHash string
)

const (
	flagConfigFilePath = "config-file-path"
)

func server() *cobra.Command {
	command := &cobra.Command{
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			configFilePath, err := cmd.Flags().GetString(flagConfigFilePath)
			if err != nil {
				return err
			}

			// init server with file config path
			app, cleanup, err := wiring.InitializeStandaloneServer(configs.ConfigFilePath(configFilePath))
			if err != nil {
				return err
			}

			defer cleanup()

			return app.Start()
		},
	}

	command.Flags().String(flagConfigFilePath, "", "If provided, will use the provided config file.")

	return command
}

func main() {
	rootCommand := &cobra.Command{
		Version: fmt.Sprintf("%s-%s", version, commitHash),
	}
	rootCommand.AddCommand(
		server(),
	)

	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
