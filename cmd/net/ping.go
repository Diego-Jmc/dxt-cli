package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var CmdPing = &cobra.Command{
	Use:   "ping",
	Short: "Ping a remote server",
	Long:  "Ping sends a message to a remote server to check if it is reachable.",
	Run: func(cmd *cobra.Command, args []string) {
		urlAddress := cmd.Flag("url").Value.String()
		if urlAddress == "" {
			fmt.Println("Please provide a URL address")
		} else {
			pingRemoteServer(urlAddress)
		}
	},
}

func init() {
	CmdPing.PersistentFlags().String("url", "www.google.com", "url where to send the ping signal")
}

func pingRemoteServer(urlAddress string) {
	formattedInstruction := fmt.Sprintf("ping %s", urlAddress)

	cmdPing := exec.Command("cmd", "/C", formattedInstruction)
	out, err := cmdPing.Output()

	if err != nil {
		fmt.Println("Invalid url address")
	} else {
		fmt.Println(string(out))
	}
}
