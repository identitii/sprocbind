package main

import (
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	var input = flag.String("input", "", "comma separated input tsql file(s) containing tables and procedures")
	var output = flag.String("output", "", "output file")
	var genTables = flag.Bool("test-gen-tables", false, "(TESTING) generate from standard non-TYPE tables too")
	var goPackage = flag.String("go-package", "", "go package to use in the generated file")

	flag.Parse()

	if *input == "" {
		log.Fatal("must set input argument")
	}

	if *output == "" {
		log.Fatal("must set output argument")
	}

	if *goPackage == "" {
		log.Fatal("must set go-package argument")
	}

	var tsql string

	for _, file := range strings.Split(*input, ",") {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("failed loading input file %s: %s", *input, err)
		}
		tsql += "\n" + string(content)
	}

	result, err := Parse(string(tsql), *genTables)
	if err != nil {
		log.Fatalf("failed parsing input file %s: %s", *input, err)
	}

	source, err := Generate(*goPackage, result)
	if err != nil {
		log.Fatalf("failed generating from input file %s: %s", *input, err)
	}

	byteSource, err := format.Source([]byte(source))
	if err != nil {
		log.Fatalf("failed formatting output source: %s", err)
	}

	err = ioutil.WriteFile(*output, byteSource, 0644)
	if err != nil {
		log.Fatalf("failed writing to output file %s: %s", *output, err)
	}

}
