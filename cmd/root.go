package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "prometheus",
	Short: "Prometheus demo server",
	Long:  `A Prometheus demo server`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
}
