package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// Badge with the number of GitHub Organizations
func organizationHandler(client github.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/organizations" {
			http.Error(w, "404 not found", http.StatusNotFound)
		}
		if r.Method != "GET" {
			http.Error(w, "method is not supported", http.StatusNotFound)
		}

		// Get GitHub username, i.e. /organizations?username=dtemir
		username := r.URL.Query().Get("username")

		// Fetch the number of Organizations the user has
		// docs.github.com/en/rest/orgs/orgs?apiVersion=2022-11-28#list-organizations-for-a-user
		orgs, _, err := client.Organizations.List(context.Background(), username, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "Number of organizations for %s is %d", username, len(orgs))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	http.HandleFunc("/organizations", organizationHandler(*client))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
