package processes

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var CmdProcess = &cobra.Command{

	Use:   "processes",
	Short: "List all the processes running in the System",
	Long:  "This function provides a comprehensive list of all processes currently running on the Windows operating system.",
	Run: func(cmd *cobra.Command, args []string) {

		cmdR := exec.Command("cmd", "/C", "tasklist")
		out, err := cmdR.Output()

		if err != nil {
			fmt.Println("Error trying to list the system processes")
		} else {
			listProcesses(string(out))
		}

	},
}

func showSimplifiedProcessDescription(process string) {
	simplified := strings.Fields(process)
	fmt.Println(fmt.Sprintf("%s  -- %s", simplified[0], simplified[1]))
}

func listProcesses(osProcesses string) {

	splitOsProcesses := strings.Split(osProcesses, "\n")
	fmt.Println("Process      PID")
	for i := 4; i < len(splitOsProcesses); i++ {
		if len(splitOsProcesses[i]) > 0 {
			showSimplifiedProcessDescription(splitOsProcesses[i])
		}
	}
}
