package main

import (
	"dxt-tools/cmd/processes"
	"fmt"
	"github.com/spf13/cobra"
)

var cmdPrintHello = &cobra.Command{
	Use:   "print",
	Short: "Print hello in console",
	Long:  "This is a  command for testing the app for the first time",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" dxt cli tool -- Developed by Diego Jiménez Chamorro")
	},
}

func main() {

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrintHello)
	rootCmd.AddCommand(processes.CmdProcess)

	rootCmd.Execute()
}
