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
	"path/filepath"
	"strings"
)

func main() {

	fs := flagset.NewFlagSet("server")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	// enable_cors := fs.Bool("enable-cors", false, "Enable CORS headers for API responses")

	enable_parse_api := fs.Bool("enable-parse-api", true, "Enable the EDTF parse API endpoint")
	enable_valid_api := fs.Bool("enable-valid-api", true, "Enable the EDTF valid API endpoint")
	enable_matches_api := fs.Bool("enable-matches-api", true, "Enable the EDTF matches API endpoint")

	enable_edtf_string_api := fs.Bool("enable-edtf-string-api", true, "Enable the SFO Museum to-edtf-string API endpoint")
	enable_edtf_date_api := fs.Bool("enable-edtf-date-api", true, "Enable the SFO Museum to-edtf-date API endpoint")

	enable_www := fs.Bool("enable-www", true, "Enable the user-facing web application.")

	path_parse_api := fs.String("path-parse-api", "/api/edtf/parse", "The path to listen for requests to the EDTF parse API on.")
	path_valid_api := fs.String("path-valid-api", "/api/edtf/valid", "The path to listen for requests to the EDTF valid API on.")
	path_matches_api := fs.String("path-matches-api", "/api/edtf/matches", "The path to listen for requests to the EDTF matches API on.")
	path_edtf_string_api := fs.String("path-edtf-string-api", "/api/sfomuseum/to-edtf-string", "The path to listen for requests to the SFO Museum to-edtf-string API on.")
	path_edtf_date_api := fs.String("path-edtf-date-api", "/api/sfomuseum/to-edtf-date", "The path to listen for requests to the SFO Museum to-edtf-date API on.")

	path_www := fs.String("path-www", "/", "The path to listen for requests to the user-facing web application on.")

	path_prefix := fs.String("path-prefix", "", "A relative path to append to all the paths the server will listen for requests on.")

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

	if *path_prefix != "" {

		log.Println("HELLO")

		*path_parse_api = filepath.Join(*path_prefix, *path_parse_api)
		*path_valid_api = filepath.Join(*path_prefix, *path_valid_api)
		*path_matches_api = filepath.Join(*path_prefix, *path_matches_api)
		*path_edtf_string_api = filepath.Join(*path_prefix, *path_edtf_string_api)
		*path_edtf_date_api = filepath.Join(*path_prefix, *path_edtf_date_api)
		*path_www = filepath.Join(*path_prefix, *path_www)
	}

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	mux := http.NewServeMux()

	if *enable_www {

		err := bootstrap.AppendAssetHandlersWithPrefix(mux, *path_prefix)

		if err != nil {
			log.Fatalf("Failed to append Bootstrap asset handlers, %v", err)
		}

		index_handler, err := sfom_www.IndexHandler()

		if err != nil {
			log.Fatalf("Failed to create WWW index handler, %v", err)
		}

		if *path_prefix != "" {

			index_handler = http.StripPrefix(*path_prefix, index_handler)

			if !strings.HasSuffix(*path_www, "/") {
				*path_www = fmt.Sprintf("%s/", *path_www)
			}
		}

		bootstrap_opts := bootstrap.DefaultBootstrapOptions()
		index_handler = bootstrap.AppendResourcesHandlerWithPrefix(index_handler, bootstrap_opts, *path_prefix)

		mux.Handle(*path_www, index_handler)
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

		mux.Handle(*path_parse_api, api_parse_handler)
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

		mux.Handle(*path_valid_api, api_valid_handler)
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

		mux.Handle(*path_matches_api, api_matches_handler)
	}

	if *enable_edtf_string_api {

		api_string_handler, err := sfom_api.ToEDTFStringHandler()

		if err != nil {
			log.Fatalf("Failed to create api.ToEDTFString handler, %v", err)
		}

		mux.Handle(*path_edtf_string_api, api_string_handler)
	}

	if *enable_edtf_date_api {

		api_date_handler, err := sfom_api.ToEDTFDateHandler()

		if err != nil {
			log.Fatalf("Failed to create api.ToEDTFDate handler, %v", err)
		}

		mux.Handle(*path_edtf_date_api, api_date_handler)
	}

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

}
