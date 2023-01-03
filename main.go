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

		values := r.URL.Query()
		// Get GitHub username, i.e. /organizations?username=dtemir
		username := values.Get("username")

		// Get shields.io color, default to brightgreen
		color := values.Get("color")
		if len(color) == 0 {
			color = "brigthgreen"
		}

		// Get shields.io style, default to empty string (shield.io's flat style)
		style := values.Get("style")

		// Get logo, default to none
		logo := values.Get("logo")

		// Fetch the number of Organizations the user has
		// docs.github.com/en/rest/orgs/orgs?apiVersion=2022-11-28#list-organizations-for-a-user
		orgs, _, err := client.Organizations.List(context.Background(), username, nil)
		if err != nil {
			log.Fatal("Error fetching organization list\n %w", username, err)
		}

		// Redirect to shields.io badge
		url := fmt.Sprintf("https://img.shields.io/badge/Organizations-%d-%s?style=%s&logo=%s", len(orgs), color, style, logo)
		http.Redirect(w, r, url, http.StatusOK)
	}
}

func main() {
	// Load in .env file with GitHub Token
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n %w", err)
	}

	// Create a GitHub client to make API calls
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Show number of organizations for user
	http.HandleFunc("/organizations", organizationHandler(*client))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server\n %w", err)
	}
}
