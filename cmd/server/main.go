package main

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-server"
	edtf_api "github.com/sfomuseum/go-edtf-http/api"
	"github.com/sfomuseum/go-flags/flagset"
	sfom_api "github.com/sfomuseum/go-sfomuseum-edtf/api"
	sfom_www "github.com/sfomuseum/go-sfomuseum-edtf/www"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	fs := flagset.NewFlagSet("server")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

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
	path_static := fs.String("path-static", "/static", "The path to listen for requests to the user-facing web application on.")

	bootstrap_prefix := fs.String("bootstrap-prefix", "", "A relative path to append to all Bootstrap-related paths the server will listen for requests on.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "HTTP server for exposing EDTF-related API methods.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "EDTF", true)

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

		// Okay, see this? The reason we've got two handlers here is that
		// if we bundle the app's static assets, specifically the JS, is
		// not handled correctly by the bootstrap.AppendResourcesHandler*
		// methods. My hunch is that the aaronland/go-http-rewrite code
		// that is used by the bootstrap code is expecting an explicit
		// content type that the io/http.FS implementations are not (?)
		// including. So for now we have separate handlers...
		// (20210111/thisisaaronland)

		static_handler, err := sfom_www.StaticHandler()

		if err != nil {
			log.Fatalf("Failed to create WWW static handler, %v", err)
		}

		path_js := filepath.Join(*path_static, "javascript/")
		path_css := filepath.Join(*path_static, "css/")
		path_wasm := filepath.Join(*path_static, "wasm/")

		path_js = fmt.Sprintf("%s/", path_js)
		path_css = fmt.Sprintf("%s/", path_css)
		path_wasm = fmt.Sprintf("%s/", path_wasm)

		mux.Handle(path_js, http.StripPrefix(*path_static, static_handler))
		mux.Handle(path_css, http.StripPrefix(*path_static, static_handler))
		mux.Handle(path_wasm, http.StripPrefix(*path_static, static_handler))

		application_handler, err := sfom_www.ApplicationHandler()

		if err != nil {
			log.Fatalf("Failed to create WWW index handler, %v", err)
		}

		bootstrap_opts := bootstrap.DefaultBootstrapOptions()
		application_handler = bootstrap.AppendResourcesHandlerWithPrefix(application_handler, bootstrap_opts, *bootstrap_prefix)

		mux.Handle(*path_www, application_handler)
	}

	if *enable_parse_api {

		api_parse_handler, err := edtf_api.ParseHandler()

		if err != nil {
			log.Fatalf("Failed to create API parse handler, %v", err)
		}

		mux.Handle(*path_parse_api, api_parse_handler)
	}

	if *enable_valid_api {

		api_valid_handler, err := edtf_api.IsValidHandler()

		if err != nil {
			log.Fatalf("Failed to create API is valid handler, %v", err)
		}

		mux.Handle(*path_valid_api, api_valid_handler)
	}

	if *enable_matches_api {

		api_matches_handler, err := edtf_api.MatchesHandler()

		if err != nil {
			log.Fatalf("Failed to create API is matches handler, %v", err)
		}

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
