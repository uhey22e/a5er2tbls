package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/uhey22e/a5er2tbls"
)

var (
	input  = flag.String("i", "", "Input a5er file.")
	output = flag.String("o", "", "Output yaml file.")
)

func handleError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func init() {
	flag.Parse()
}

func main() {
	r, err := a5er2tbls.ParseRelations(*input)
	handleError(err)

	y, err := yaml.Marshal(r)
	handleError(err)

	var out *os.File
	if len(*output) > 0 {
		f, err := os.Create(*output)
		handleError(err)
		out = f
	} else {
		out = os.Stdout
	}
	defer out.Close()

	_, err = out.Write(y)
	handleError(err)
}
