package main

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-server"
	// "github.com/rs/cors"
	edtf_api "github.com/sfomuseum/go-edtf-http/api"
	"github.com/sfomuseum/go-flags/flagset"
	sfom_api "github.com/sfomuseum/go-sfomuseum-edtf/api"
	sfom_www "github.com/sfomuseum/go-sfomuseum-edtf/www"
	"log"
	"net/http"
	"os"
)

func main() {

	fs := flagset.NewFlagSet("server")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	// enable_cors := fs.Bool("enable-cors", false, "Enable CORS headers for API responses")

	enable_parse_api := fs.Bool("enable-parse-api", true, "Enable the /api/edtf/parse endpoint")
	enable_valid_api := fs.Bool("enable-valid-api", true, "Enable the /api/edtf/valid endpoint")
	enable_matches_api := fs.Bool("enable-matches-api", true, "Enable the /api/edtf/matches endpoint")

	enable_edtf_string_api := fs.Bool("enable-edtf-string-api", true, "Enable the /api/sfomuseum/to-edtf-string endpoint")
	enable_edtf_date_api := fs.Bool("enable-edtf-date-api", true, "Enable the /api/sfomuseum/to-edtf-date endpoint")

	enable_www := fs.Bool("enable-www", true, "Enable the user-facing web interface")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "HTTP server for exposing EDTF-related API methods.\n")
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

	if *enable_www {

		err := bootstrap.AppendAssetHandlers(mux)

		if err != nil {
			log.Fatalf("Failed to append Bootstrap asset handlers, %v", err)
		}

		index_handler, err := sfom_www.IndexHandler()

		if err != nil {
			log.Fatalf("Failed to create WWW index handler, %v", err)
		}

		bootstrap_opts := bootstrap.DefaultBootstrapOptions()
		index_handler = bootstrap.AppendResourcesHandler(index_handler, bootstrap_opts)

		mux.Handle("/", index_handler)
	}

	if *enable_parse_api {

		api_parse_handler, err := edtf_api.ParseHandler()

		if err != nil {
			log.Fatalf("Failed to create API parse handler, %v", err)
		}

		/*
			if *enable_cors {
			api_parse_handler = c.Handler(api_parse_handler)
			}
		*/

		mux.Handle("/api/edtf/parse", api_parse_handler)
	}

	if *enable_valid_api {

		api_valid_handler, err := edtf_api.IsValidHandler()

		if err != nil {
			log.Fatalf("Failed to create API is valid handler, %v", err)
		}

		/*
			if *enable_cors {
				api_valid_handler = c.Handler(api_valid_handler)
			}
		*/

		mux.Handle("/api/edtf/valid", api_valid_handler)
	}

	if *enable_matches_api {

		api_matches_handler, err := edtf_api.MatchesHandler()

		if err != nil {
			log.Fatalf("Failed to create API is matches handler, %v", err)
		}

		/*
			if *enable_cors {
				api_matches_handler = c.Handler(api_matches_handler)
			}
		*/

		mux.Handle("/api/edtf/matches", api_matches_handler)
	}

	if *enable_edtf_string_api {

		api_string_handler, err := sfom_api.ToEDTFStringHandler()

		if err != nil {
			log.Fatalf("Failed to create api.ToEDTFString handler, %v", err)
		}

		mux.Handle("/api/sfomuseum/to-edtf-string", api_string_handler)
	}

	if *enable_edtf_date_api {

		api_date_handler, err := sfom_api.ToEDTFDateHandler()

		if err != nil {
			log.Fatalf("Failed to create api.ToEDTFDate handler, %v", err)
		}

		mux.Handle("/api/sfomuseum/to-edtf-date", api_date_handler)
	}

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

}
