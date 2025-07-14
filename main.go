package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"net/http"
	"strings"
)

var (
	Url string
)

func main() {
	pflag.StringVar(&Url, "url", "https://google.com", "fill in values")
	pflag.Parse()

	resp, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Protocol: %s\n", resp.Proto)
	fmt.Print("Header:")
	for key, values := range resp.Header {
		fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
	}
	defer resp.Body.Close()

}
