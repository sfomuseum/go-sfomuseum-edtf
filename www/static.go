package www

import (
	"embed"
	"net/http"
)

// See note in cmd/server/main.go for why this is distinct
// from index.go (20210111/thisisaaronland)

//go:embed javascript
//go:embed css
var static_fs embed.FS

func StaticHandler() (http.Handler, error) {

	http_fs := http.FS(static_fs)
	fs_handler := http.FileServer(http_fs)

	return fs_handler, nil
}
