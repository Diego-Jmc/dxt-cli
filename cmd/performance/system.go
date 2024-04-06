package performance

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
)

var CmdSystem = &cobra.Command{
	Use:   "system",
	Short: "Show system info",
	Long:  "Provide information about the system including ,hardware and CPU usage.",
	Run: func(cmd *cobra.Command, args []string) {
		showHardwareMemory()
	},
}

func getMbFromBytes(bytes uint64) uint64 {
	return (bytes / 1000000)
}

func getGbFromMegabytes(megabytes uint64) uint64 {
	return megabytes / 1024
}

func getGbFromBytes(bytes uint64) uint64 {
	return bytes / (1024 * 1024 * 1024)

}

func showHardwareMemory() {

	var freeBytesAvailable uint64
	var totalNumberOfBytes uint64
	var totalNumberOfFreeBytes uint64

	err := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr("C:"),
		&freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Total space: %d GB\n", getGbFromBytes(totalNumberOfBytes))
		fmt.Printf("Available space : %d GB\n", getGbFromBytes(freeBytesAvailable))
	}

}
