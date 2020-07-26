package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/uhey22e/a5er2tbls"
)

var (
	input       = flag.String("i", "", "Input a5er file.")
	output      = flag.String("o", "", "Output yaml file.")
	showVersion = flag.Bool("v", false, "Show application version.")
)

// These values are set by build arguments.
var (
	version = "unknown"
)

func handleError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func writeOut(b []byte) error {
	var out *os.File
	if len(*output) > 0 {
		f, err := os.Create(*output)
		if err != nil {
			return err
		}
		out = f
	} else {
		out = os.Stdout
	}
	defer out.Close()

	_, err := out.Write(b)
	return err
}

func init() {
	flag.Parse()
}

func main() {
	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	r, err := a5er2tbls.ParseRelations(*input)
	handleError(err)

	y, err := yaml.Marshal(r)
	handleError(err)

	err = writeOut(y)
	handleError(err)
}
