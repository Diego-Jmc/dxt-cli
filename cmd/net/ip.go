package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"net/http"
)

var Cmdip = &cobra.Command{
	Use:   "ip",
	Short: "Show public and private ip address",
	Run: func(cmd *cobra.Command, args []string) {
		showIpAddressDetails()
	},
}

func showIpAddressDetails() {
	showPublicIpAddress()
	showPrivateIpAddress()
}

func showPublicIpAddress() {
	const url = "https://api.ipify.org?format=text" // we are using a public IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("your public IP is:%s\n", ip)
}

func showPrivateIpAddress() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	fmt.Printf("your private IP is:%s\n", localAddress.IP.String())
}

func init() {

}
