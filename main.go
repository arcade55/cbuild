package main

import (
	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	projectID := "k7endpoint"

	// Creates a trace client.
	traceClient, err := trace.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	p, err := trace.NewLimitedSampler(0.999, 5)
	traceClient.SetSamplingPolicy(p)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span := traceClient.NewSpan("newTraceSpanForBuild")
		defer span.Finish()
//		time.Sleep(1120 * time.Millisecond)
		io.WriteString(w, "Automated build and deploy!!!!!")
	})

	http.Handle("/foo", traceClient.HTTPHandler(handler))
	log.Fatal(http.ListenAndServe(":6060", nil))
}
