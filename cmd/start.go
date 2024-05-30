package cmd

import (
	"github.com/parvejmia9/Prometheus/api"
	"github.com/spf13/cobra"
)

var Port string

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			api.RunServer(Port)
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&Port, "port", "p", "8181", "Port to listen on")
	RootCmd.AddCommand(StartCmd)
}
