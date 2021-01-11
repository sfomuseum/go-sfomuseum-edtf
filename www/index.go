package www

import (
	"net/http"
	"embed"
)

//go:embed index.html
//go:embed javascript
//go:embed css
var web_app embed.FS

func IndexHandler() (http.HandlerFunc, error) {

	http_fs := http.FS(web_app)
	fs_handler := http.FileServer(http_fs)

	return fs_handler, nil
}
