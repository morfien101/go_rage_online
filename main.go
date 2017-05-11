// This application simply burns CPU cycles in an attempt to stress a system.
// Designed to allow for testing of auto-scalling groups where triggers are
// required.
// To use it simply run it and curl it on port 8001

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	showVersion = flag.Bool("v", false, "Shows the version of the application")
	httpPort    = flag.Int("p", 8001, "The port number to run the http server on")
	httpAddress = flag.String("a", "0.0.0.0", "The address to run the HTTP server on.")
	help        = flag.Bool("h", false, "Shows the help menu.")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	startHTTPServer()
}

func cpuStress() {
	var number string
	number = ""
	fmt.Println("CPU Stress start: ", time.Now())
	rand.Seed(int64(time.Now().Nanosecond()))
	loopindex := rand.Intn(500000)
	fmt.Println("looping", loopindex, "times")
	for i := 0; i < loopindex; i++ {
		number += string(i)
	}
	fmt.Println("CPU Stress End: ", time.Now())
}
