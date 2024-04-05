package processes

import (
	"github.com/spf13/cobra"
)

var cmdProcess = &cobra.Command{
	Use:   "processes",
	Short: "List all the processes running",
	Long:  "This function provides a comprehensive list of all processes currently running on the Windows operating system.",
}
