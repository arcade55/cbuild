package main

import (
	"log"
	"io"
	"net/http"
	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
)

func main() {

	ctx := context.Background()
	projectID := "k7endpoint"

	// Creates a trace client.
	traceClient, err := trace.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "With tracing capability???")
	})

	http.Handle("/foo", traceClient.HTTPHandler(handler))
	log.Fatal(http.ListenAndServe(":6060", nil))
}
