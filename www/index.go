package www

import (
	"embed"
	"net/http"
)

// See note in cmd/server/main.go for why this is distinct
// from static.go (20210111/thisisaaronland)

//go:embed index.html
var web_app embed.FS

func IndexHandler() (http.Handler, error) {

	http_fs := http.FS(web_app)
	fs_handler := http.FileServer(http_fs)

	return fs_handler, nil
}
