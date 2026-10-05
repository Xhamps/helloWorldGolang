package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed index.html
var fs embed.FS

var tmpl = template.Must(template.ParseFS(fs, "index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		tmpl.Execute(w, nil)
	})
	// htmx target: returns a fragment, not a full page
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<h1 class="glow">HELLO, WORLD</h1><p class="sub">transmission received · ` + time.Now().Format("15:04:05") + `</p>`))
	})
	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
