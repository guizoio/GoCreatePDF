package main

import (
	"CreateFilePDF/src/cmd"
	"CreateFilePDF/src/configs"
	"CreateFilePDF/src/infra"
	"context"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

func init() {
	configs.LoadEnv()
}

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	containerDI := infra.NewContainerDI()
	defer containerDI.ShutDown()

	cmdMakeMigrations := &cobra.Command{
		Use:   "MakeMigrations",
		Short: "Run MakeMigrations",
		Run: func(cli *cobra.Command, args []string) {
			makeMigration := cmd.NewDatabaseMakeMigrations(containerDI.DB)
			makeMigration.MakeMigrations()
		},
	}

	cmdHttpServer := &cobra.Command{
		Use:   "httpserver",
		Short: "Run httpserver",
		Run: func(cli *cobra.Command, args []string) {
			cmd.StartHttp(ctx, containerDI)
		},
	}

	var rootCmd = &cobra.Command{Use: "APP"}
	rootCmd.AddCommand(cmdMakeMigrations)
	rootCmd.AddCommand(cmdHttpServer)
	rootCmd.Execute()

}
