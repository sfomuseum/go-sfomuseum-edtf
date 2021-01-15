package www

import (
	"embed"
	"net/http"
	"log"
	"io/fs"
)

// See note in cmd/server/main.go for why this is distinct
// from static.go (20210111/thisisaaronland)

//go:embed index.html
//go:embed validate
var web_app embed.FS

func ApplicationHandler() (http.Handler, error) {

	fn := func(path string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			return nil
		}

		log.Println(path)
		return nil
	}

	fs.WalkDir(web_app, ".", fn)
	
	http_fs := http.FS(web_app)
	fs_handler := http.FileServer(http_fs)

	return fs_handler, nil
}
