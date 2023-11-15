package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sfomuseum/go-sfomuseum-edtf"	
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more SFO Museum date strings and return a list of JSON-encode edtf.EDTFDate objects.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s date(N) date(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	writers := []io.Writer{
		os.Stdout,
	}

	wr := io.MultiWriter(writers...)

	wr.Write([]byte(`[`))

	for i, sfom_str := range flag.Args() {

		d, err := edtf.ToEDTFDate(sfom_str)

		if err != nil {
			log.Fatalf("Failed to parse EDTF string '%s', %v", sfom_str, err)
		}

		enc, err := json.Marshal(d)

		if err != nil {
			log.Fatalf("Failed to encode EDTFDate, %v", err)
		}

		if i > 0 {
			wr.Write([]byte(`,`))
		}

		wr.Write(enc)
	}

	wr.Write([]byte(`]`))
}
