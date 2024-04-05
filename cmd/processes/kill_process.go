package processes

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"syscall"
)

var cmdKillProcess = &cobra.Command{
	Use:   "kill",
	Short: "Kill a process by PID",
	Long:  "Kill a system process given the PID",
	Run: func(cmd *cobra.Command, args []string) {

		pidFlag := cmd.Flag("pid").Value.String()
		pid, err := strconv.Atoi(pidFlag)

		if err != nil {
			fmt.Println("PID is not a number")
		} else {
			killProcess(pid)
			fmt.Sprintf("Process with pid: %d killed", pid)
		}

	},
}

func init() {
	cmdKillProcess.PersistentFlags().Int("pid", 0, "Kill a task by pid")
}

func killProcess(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Print("Could not kill process : Process not found")
	} else {
		fmt.Print("Could not kill process : Process not found")
		process.Signal(syscall.SIGTERM)
	}
}
