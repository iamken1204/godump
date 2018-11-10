package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godump",
	Short: "godump is mean to be a portable tcpdump alternative",
	Long:  "godump is mean to be a portable tcpdump alternative",
	Run: func(cmd *cobra.Command, args []string) {
		initMessage()
		tracePacket()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("godump panics:", err)
	}
}

func initMessage() {
	fmt.Printf(`godump is monitoring
Device:   %s
Filter:   %s
Show Plain Data:          %t
Show Application Payload: %t

`, device, tcpdumpFilter, dump, app)
}
