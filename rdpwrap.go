package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := r.URL.Path[1:] // skip leading slash
		fullpath := "data/" + filepath

		f, err := os.OpenFile(fullpath, os.O_RDONLY, 0644)
		if err != nil {
			http.Error(w, fmt.Sprintf("404 找不到文件 %q: %v", fullpath, err), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// serve the file
		http.ServeContent(w, r, filepath, time.Now(), f)
	})

	http.ListenAndServe(":8080", nil)
}
