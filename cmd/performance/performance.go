package performance

import "github.com/spf13/cobra"

var CmdProcess = &cobra.Command{
	Use:   "performance",
	Short: "System performance summary",
	Long: `The performance command provides a summary of system performance metrics
			such as CPU usage, memory utilization, disk I/O, and network activity`,
}

func init() {
	CmdProcess.AddCommand(CmdSystem)
}
