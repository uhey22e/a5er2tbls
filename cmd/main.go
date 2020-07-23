package main

import (
	"flag"
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/uhey22e/a5er2tbls"
)

var (
	input  = flag.String("i", "", "Input a5er file.")
	output = flag.String("o", "", "Output yaml file.")
)

func main() {
	flag.Parse()
	r := a5er2tbls.ParseRelations(*input)
	y, _ := yaml.Marshal(r)
	fmt.Println(string(y))
}
