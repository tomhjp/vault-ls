package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	srv := http.FileServer(http.Dir("."))
	fmt.Println("Serving on localhost:8080")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		srv.ServeHTTP(resp, req)
	})); err != nil {
		_, _ = fmt.Println("Error shutting down server", err)
		os.Exit(1)
	}
}
