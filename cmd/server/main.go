package main

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-server"
	// "github.com/rs/cors"
	edtf_api "github.com/sfomuseum/go-edtf-http/api"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"net/http"
	"os"
)

func main() {

	fs := flagset.NewFlagSet("server")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	// enable_cors := fs.Bool("enable-cors", false, "Enable CORS headers for API responses")

	enable_parse_api := fs.Bool("enable-parse-api", true, "Enable the /api/parse endpoint")
	enable_valid_api := fs.Bool("enable-valid-api", true, "Enable the /api/valid endpoint")
	enable_matches_api := fs.Bool("enable-matches-api", true, "Enable the /api/matches endpoint")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "HTTP server for exposing sfomuseum/go-edtf-http handlers.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "EDTF")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	mux := http.NewServeMux()

	if *enable_parse_api {

		api_parse_handler, err := edtf_api.ParseHandler()

		if err != nil {
			log.Fatalf("Failed to API parse handler, %v", err)
		}

		/*
			if *enable_cors {
			api_parse_handler = c.Handler(api_parse_handler)
			}
		*/

		mux.Handle("/api/parse", api_parse_handler)
	}

	if *enable_valid_api {

		api_valid_handler, err := edtf_api.IsValidHandler()

		if err != nil {
			log.Fatalf("Failed to API is valid handler, %v", err)
		}

		/*
			if *enable_cors {
				api_valid_handler = c.Handler(api_valid_handler)
			}
		*/

		mux.Handle("/api/valid", api_valid_handler)
	}

	if *enable_matches_api {

		api_matches_handler, err := edtf_api.MatchesHandler()

		if err != nil {
			log.Fatalf("Failed to API is matches handler, %v", err)
		}

		/*
			if *enable_cors {
				api_matches_handler = c.Handler(api_matches_handler)
			}
		*/

		mux.Handle("/api/matches", api_matches_handler)
	}

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

}
