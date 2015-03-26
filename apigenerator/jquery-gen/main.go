package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ericaro/jquery/apigenerator"
)

var (
	output = flag.String("o", "", "output directory (default to os.Stdout)")
	input  = flag.String("i", "", "input directory where are the entries.xml ")
)

func main() {
	flag.Parse()

	// create a writer either file (-o option) or stdout
	var target io.Writer
	if *output == "" {
		target = os.Stdout
	} else {
		file, err := os.Create(*output)
		if err != nil {
			panic(fmt.Errorf("cannot write to %v: %v", output, err))
		}
		defer file.Close()
		target = file
	}

	// parse each entry as described in th
	entries, err := apigenerator.Parse(*input)
	if err != nil {
		fmt.Printf("Error parsing xml entries: %v\n", err)
		os.Exit(-1)
	}

	apigenerator.Format(target, entries...)
}
