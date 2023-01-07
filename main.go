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
	gh_client := github.NewClient(tc)

	// Create a MongoDB client to store visit counts

	// Show number of organizations for user
	http.HandleFunc("/organizations", organizationsHandler(*gh_client))

	// Show number of years user has been a GitHub member
	http.HandleFunc("/years", yearsHandler(*gh_client))

	// Show number of repos user owns
	http.HandleFunc("/repos", reposHandler(*gh_client))

	// Show number of visits to user's GitHub page
	http.HandleFunc("/visits", visitsHandler(*gh_client))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server\n %w", err)
	}
}
