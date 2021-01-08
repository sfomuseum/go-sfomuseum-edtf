package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-edtf"
	"log"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more SFO Museum date strings and return a line-separated list of valid EDTF strings.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s date(N) date(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, sfom_str := range flag.Args() {

		edtf_str, err := edtf.ToEDTFString(sfom_str)

		if err != nil {
			log.Fatalf("Failed to parse '%s', %v", sfom_str, err)
		}

		fmt.Println(edtf_str)
	}
}
