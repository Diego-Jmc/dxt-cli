package main

import (
	"dxt-tools/cmd/net"
	"dxt-tools/cmd/performance"
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

	var rootCmd = &cobra.Command{Use: "dxt"}
	rootCmd.AddCommand(cmdPrintHello)
	rootCmd.AddCommand(processes.CmdProcess)
	rootCmd.AddCommand(net.CmdProcess)
	rootCmd.AddCommand(performance.CmdProcess)

	rootCmd.Execute()

}
