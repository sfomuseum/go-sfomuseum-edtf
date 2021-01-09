package api

import (
	"encoding/json"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-edtf/parser"
	"net/http"
)

type MatchesResult struct {
	Level   int    `json:"level"`
	Feature string `json:"feature"`
}

func MatchesHandler() (http.HandlerFunc, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		edtf_str, err := sanitize.GetString(req, "edtf")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		if edtf_str == "" {
			http.Error(rsp, "Empty '?edtf=' parameter", http.StatusBadRequest)
			return
		}

		level, feature, err := parser.Matches(edtf_str)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		r := MatchesResult{
			Level:   level,
			Feature: feature,
		}

		rsp.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(r)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
