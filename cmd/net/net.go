package net

import (
	"github.com/spf13/cobra"
)

var CmdProcess = &cobra.Command{
	Use:   "net",
	Short: "Set of net operations",
	Long:  "net is a command that provides an abstraction for net operations",
}

func init() {
	CmdProcess.AddCommand(Cmdip)
	CmdProcess.AddCommand(CmdPing)
}
