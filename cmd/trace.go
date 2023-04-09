/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP address",
	Long: `Trace the Ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		if(len(args) == 0) {
			fmt.Println("Please provide an IP address")
		}else{
			for _, ip := range args {
				showData(ip)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct{
	IP string `json::"ip"`
	City string `json::"city"`
	Region string `json::"region"`
	Country string `json::"country"`
	Loc string `json::"loc"`
	Timezone string `json::"timezone"`
}


func showData(ip string){
	responseByte:= getData(ip)
	data := Ip{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println("Unable to unmarshal the response")
	}
	c := color.New(color.FgBlue).Add(color.Bold)
	c.Println("Data found:")
	fmt.Println("IP: ", data.IP)
	fmt.Println("City: ", data.City)
	fmt.Println("Region: ", data.Region)
	fmt.Println("Country: ", data.Country)
	fmt.Println("Location: ", data.Loc)
	fmt.Println("Timezone: ", data.Timezone)
	fmt.Println()
}

func getData(ip string) []byte {
	url := "https://ipinfo.io/" + ip + "/geo"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Unable to get the response")
		os.Exit(1)
	}
	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read the response")
		os.Exit(1)
	}

	return responseByte 
}


