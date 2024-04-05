package processes

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strconv"
	"strings"
)

var cmdListProcesses = &cobra.Command{
	Use:   "ls",
	Short: "List processes",
	Long:  "This function provides a comprehensive list of all processes currently running on the Windows operating system.",
	Run: func(cmd *cobra.Command, args []string) {
		amount := cmd.Flag("amount").Value.String()

		cmdR := exec.Command("cmd", "/C", "tasklist")
		out, err := cmdR.Output()

		if err == nil {

			if amount == "all" {

				// TODO: Solve why we cant use listProcesses()
				listProcessesAll(string(out))

			} else {

				intAmount, err := strconv.Atoi(amount)

				if err != nil {
					fmt.Println("amount must be an integer")
					return
				}

				if intAmount < 0 {
					fmt.Println("amount must be a positive integer")
					return
				} else {
					listProcesses(string(out), intAmount)
				}
			}
		} else {
			fmt.Println("Error trying to read processes")
		}
	},
}

func showSimplifiedProcessDescription(process string) {
	simplified := strings.Fields(process)
	fmt.Println(fmt.Sprintf("%s  -- %s", simplified[0], simplified[1]))
}

func listProcessesAll(osProcesses string) {

	splitOsProcesses := strings.Split(osProcesses, "\n")
	fmt.Println("Process      PID")

	for i := 4; i < len(splitOsProcesses); i++ {
		if len(splitOsProcesses[i]) >= 1 {
			showSimplifiedProcessDescription(splitOsProcesses[i])
		}
	}

}

func init() {
	CmdProcess.PersistentFlags().String("amount", "all", "Amount of processes that the user wants to see")
}

func listProcesses(osProcesses string, leght int) {

	splitOsProcesses := strings.Split(osProcesses, "\n")
	fmt.Println("Process      PID")
	for i := 4; i < leght+4; i++ {
		if len(splitOsProcesses[i]) >= 1 {
			showSimplifiedProcessDescription(splitOsProcesses[i])
		}
	}
}
