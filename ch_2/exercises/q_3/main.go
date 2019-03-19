package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	gw := gzip.NewWriter(w)
	writers := io.MultiWriter(gw, os.Stdout)

	encoder := json.NewEncoder(writers)
	encoder.SetIndent("", "\t")
	encoder.Encode(source)

	gw.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
