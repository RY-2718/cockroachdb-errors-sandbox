package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cockroachdb/errors"

	"github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/trace-error", handler.TraceErrorHandler)
	mux.HandleFunc("/trace-library-error", handler.TraceLibraryErrorHandler)

	fmt.Print("Server is running on http://localhost:8888\n\n")
	if err := http.ListenAndServe(":8888", mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed to start: %v", err)
	}
}
