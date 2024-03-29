package api

import (
	"encoding/json"
	"net/http"

	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-sfomuseum-edtf"	
)

func ToEDTFDateHandler() (http.HandlerFunc, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		date_str, err := sanitize.GetString(req, "date")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		if date_str == "" {
			http.Error(rsp, "Empty '?date=' parameter", http.StatusBadRequest)
			return
		}

		edtf_date, err := edtf.ToEDTFDate(date_str)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(edtf_date)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
