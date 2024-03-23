package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	webview "github.com/webview/webview_go"
)

//go:embed dist/*
var staticFiles embed.FS

func Serve() {
	staticSubFS, err := fs.Sub(staticFiles, "dist")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(staticSubFS)))

	fmt.Println("Server listening on port 8080...")
	go http.ListenAndServe(":8080", nil)

	w := webview.New(true)
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:8080")
	w.Run()
}
