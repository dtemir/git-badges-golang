package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v48/github"
)

// Badge with the number of GitHub Organizations
func organizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/organizations" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}

	// Get GitHub username, i.e. /organizations?username=dtemir
	username := r.URL.Query().Get("username")

	client := github.NewClient(nil)

	// Fetch the number of Organizations the user has
	// docs.github.com/en/rest/orgs/orgs?apiVersion=2022-11-28#list-organizations-for-a-user
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Number of organizations for %s is %d", username, len(orgs))
}

func main() {
	http.HandleFunc("/organizations", organizationHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
