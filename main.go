package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	tcp "github.com/tevino/tcp-shaker"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
	cyan  = color.New(color.FgCyan).SprintFunc()
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("\"tcping\" requires 1 argument: host:port or 2 arguments: host and port.")
		os.Exit(1)
	}

	host := args[0]
	if host == "" {
		fmt.Println("\"tcping\" error: host not found.")
		os.Exit(1)
	}
	url := host

	if !strings.Contains(host, ":") && len(args) == 2 {
		port := args[1]
		if port == "" {
			fmt.Println("\"tcping\" error: port not found.")
			os.Exit(1)
		}
		url = host + ":" + port
	}

	c := tcp.NewChecker(true)
	if err := c.InitChecker(); err != nil {
		log.Fatal("tcping init failed:", err)
	}

	timeout := time.Second * 1
	err := c.CheckAddr(url, timeout)
	switch err {
	case tcp.ErrTimeout:
		fmt.Printf(red("KO") + " Connect to " + cyan(url) + " timed out\n")
	case nil:
		fmt.Printf(green("OK") + " Connect to " + cyan(url) + " succeeded\n")
	default:
		if e, ok := err.(*tcp.ErrConnect); ok {
			fmt.Printf(red("KO")+" Connect to "+cyan(url)+" failed: %s\n", e)
		} else {
			fmt.Println(red("KO")+" Error occurred while connecting to  "+cyan(url)+":", err)
		}
	}
}
