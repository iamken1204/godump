package cmd

import (
	"time"

	"github.com/google/gopacket/pcap"
)

// godump cmd flag variables
var (
	device        string
	tcpdumpFilter string
	dump          bool
	app           bool
)

// fixed tcpdump variables
var (
	snapshopLen int32         = 1024
	promiscuous bool          = false
	timeout     time.Duration = pcap.BlockForever
)

func init() {
	rootCmd.Flags().StringVarP(&device, "interface", "i", "eth0", "internet interface card")
	rootCmd.Flags().StringVarP(&tcpdumpFilter, "filter", "f", "tcp", "tcpdump filter")
	rootCmd.Flags().BoolVarP(&dump, "dump", "d", false, "dump plain data")
	rootCmd.Flags().BoolVarP(&app, "app", "a", false, "show application layer payload")
}
