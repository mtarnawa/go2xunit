package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	// Version is the current version
	Version = "2.0.0"
)

func TestTime(input *File) time.Time {
	stat, err := input.Stat()
	if err != nil {
		return time.Now()
	}
	return stat.ModTime()
}

func main() {
	if args.showVersion {
		fmt.Printf("go2xunit %s\n", Version)
		os.Exit(0)
	}

	// No time ... prefix for error messages
	log.SetFlags(0)

	flag.Parse()
	if err := validateArgs(); err != nil {
		log.Fatalf("error: %s", err)
	}

	input, output, err := getIO(args.inFile, args.outFile)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// We'd like the test time to be the time of the generated file
	testTime := TestTime()

	xmlTemplate := XUnitTemplate
	if args.xunitnetOut {
		xmlTemplate = XUnitNetTemplate
	} else if args.bambooOut || (len(suites) > 1) {
		xmlTemplate = XMLMultiTemplate
	}

	lib.WriteXML(suites, output, xmlTemplate, testTime)
	if args.fail && suites.HasFailures() {
		os.Exit(1)
	}
}
