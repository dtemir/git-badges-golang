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
	client := github.NewClient(tc)

	// Show number of organizations for user
	http.HandleFunc("/organizations", organizationsHandler(*client))

	// Show number of years user has been a GitHub member
	http.HandleFunc("/years", yearsHandler(*client))

	// Show number of repos user owns
	http.HandleFunc("/repos", reposHandler(*client))

	// Show number of visits to user's GitHub page
	http.HandleFunc("/visits", visitsHandler(*client))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server\n %w", err)
	}
}
