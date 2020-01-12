package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/whoyare", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, _ := json.Marshal(map[string]string{
			"addr": r.RemoteAddr,
		})
		w.Write(body)
	})
	http.ListenAndServe(":4080", nil)
}
