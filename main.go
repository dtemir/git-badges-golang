package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v48/github"
)

// Return first GitHub organization
func organizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/organizations" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}

	client := github.NewClient(nil)

	orgs, _, err := client.Organizations.List(context.Background(), "dtemir", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%v", orgs[0])
}

func main() {
	http.HandleFunc("/organizations", organizationHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
