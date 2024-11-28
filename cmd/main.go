package cmd

import (
	"github.com/rezairfanwijaya/go-autoscale.git/app"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cli := &cobra.Command{}

	cli.AddCommand(&cobra.Command{
		Use:   "start-app",
		Short: "start the main app",
		Long:  "start the main rest api app",
		Run: func(cmd *cobra.Command, args []string) {
			app.StartApp()
		},
	})

	cli.AddCommand(&cobra.Command{
		Use:   "start-worker",
		Short: "start the main app",
		Long:  "start the main rest api app",
		Run: func(cmd *cobra.Command, args []string) {
			app.StartWorker()
		},
	})

	return cli
}
