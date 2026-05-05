package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		fs.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
