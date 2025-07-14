package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lnvn/web-header-crawler/v2/pkg/logger"
	"github.com/spf13/pflag"
)

var (
	Url string
)

func main() {
	log, err := logger.New()
	if err != nil {
		log.Fatal("Fail to initiate:")
	}

	defer log.Sync()

	pflag.StringVar(&Url, "url", "https://google.com", "fill in values")
	pflag.Parse()

	if !pflag.CommandLine.Changed("url") {
		fmt.Printf("--url flag not specified. Using default %s\n\n", Url)
		fmt.Print("---\n\n")
	}

	resp, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// log.Info("Printing out status, protocol & Header")
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Protocol: %s\n", resp.Proto)
	fmt.Print("Header:")
	for key, values := range resp.Header {
		fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
	}
	defer resp.Body.Close()

}
