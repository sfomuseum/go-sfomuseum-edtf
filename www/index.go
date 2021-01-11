package www

import (
	"embed"
	"net/http"
)

//go:embed index.html
//go:embed javascript
//go:embed css
var web_app embed.FS

func IndexHandler() (http.Handler, error) {

	http_fs := http.FS(web_app)
	fs_handler := http.FileServer(http_fs)

	return fs_handler, nil
}
