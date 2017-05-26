package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const version = "0.0.1"
var duration = flag.String("duration", "25m", "duration")
var showVersion = flag.Bool("version", false, "print version and exit")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		os.Exit(2)
	}
	d, err := time.ParseDuration(*duration)
	if err != nil {
		log.Fatalf("could not parse given duration %v: %v", *duration, err)
	}
	end := time.Now().Add(d)
	for {
		secs := time.Until(end).Seconds()
		if secs <= 0 {
			break
		}
		print("\033[H\033[2J")
		fmt.Printf("%02d:%02d", int(secs / 60), int(secs) % 60)
		time.Sleep(1 * time.Second)
	}
}
