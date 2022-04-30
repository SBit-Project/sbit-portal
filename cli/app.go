package cli

import (
	"log"
	"net/url"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var app = kingpin.New("sbitportal", "SBIT DApp Server")

var sbitRPC = app.Flag("sbit-rpc", "URL of sbit RPC service").Envar("SBIT_RPC").Default("").String()

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func getSBitRPCURL() *url.URL {
	if *sbitRPC == "" {
		log.Fatalln("Please set SBIT_RPC to sbitd's RPC URL")
	}

	url, err := url.Parse(*sbitRPC)
	if err != nil {
		log.Fatalln("SBIT_RPC URL:", *sbitRPC)
	}

	if url.User == nil {
		log.Fatalln("SBIT_RPC URL (must specify user & password):", *sbitRPC)
	}

	return url
}
