package processes

import (
	"github.com/spf13/cobra"
)

var CmdProcess = &cobra.Command{

	Use:   "pro",
	Short: "Command to manage processes operations",
	Long:  "Provide convenient tools for performing operations on processes",
}

func init() {
	// ######   Commands  ######
	CmdProcess.AddCommand(cmdKillProcess)
	CmdProcess.AddCommand(cmdListProcesses)

}
