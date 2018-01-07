// Command line parsing
package main

import (
	"flag"
	"fmt"
	"strings"
)

var args struct {
	inFile      string
	outFile     string
	fail        bool
	failOnRace  bool
	showVersion bool
	outType     string
	//suitePrefix string
}

var outTypes = []string{
	"jenkins",
	"xunitnet",
	"bamboo",
}

func init() {
	flag.StringVar(&args.inFile, "input", "", "input file (default to stdin)")
	flag.StringVar(&args.outFile, "output", "", "output file (default to stdout)")
	flag.BoolVar(&args.fail, "fail", false, "fail (non zero exit) if any test failed")
	flag.BoolVar(&args.showVersion, "version", false, "print version and exit")
	flag.StringVar(&args.outType, "type", "jenkins",
		fmt.Sprintf("output type (%s)", strings.Join(outTypes, ", ")))
	flag.BoolVar(&args.failOnRace, "fail-on-race", false,
		"mark test as failing if it exposes a data race")
	/*
		flag.StringVar(&args.suitePrefix, "suite-name-prefix", "",
			"prefix to include before all suite names")
	*/

	flag.Parse()
}

func inList(val string, list []string) bool {
	for _, v := range list {
		if val == v {
			return true
		}
	}
	return false
}

// validateArgs validates command line arguments
func validateArgs() error {
	if flag.NArg() > 0 {
		return fmt.Errorf("too many arguments (did you mean -input?)")
	}

	if !inList(args.outType, outTypes) {
		return fmt.Errorf("unknown output format - %q", args.outType)
	}
	return nil
}
