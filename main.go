package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// Badge with the number of your GitHub Organizations
func organizationsHandler(client github.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/organizations" {
			http.Error(w, "404 not found", http.StatusNotFound)
		}
		if r.Method != "GET" {
			http.Error(w, "Method is not supported", http.StatusNotFound)
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

		// Create a shields.io badge
		url := fmt.Sprintf("https://img.shields.io/badge/Organizations-%d-%s?style=%s&logo=%s", len(orgs), color, style, logo)

		svg := getSVG(url)

		// Display SVG badge
		w.Header().Set("content-type", "image/svg+xml")
		w.Header().Set("cache-control", "no-cache")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", svg)
	}
}

// Badge with the number of years you have been a GitHub member
func yearsHandler(client github.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/years" {
			http.Error(w, "404 not found", http.StatusNotFound)
		}
		if r.Method != "GET" {
			http.Error(w, "Method is not supported", http.StatusNotFound)
		}

		values := r.URL.Query()
		username := values.Get("username")

		color := values.Get("color")
		if len(color) == 0 {
			color = "brigthgreen"
		}

		style := values.Get("style")

		logo := values.Get("logo")

		// Fetch the user data
		// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user
		user, _, err := client.Users.Get(context.Background(), username)
		if err != nil {
			log.Fatal("Error fetching user\n %w", username, err)
		}

		// Calculate the number of years passed since user creation
		created := user.GetCreatedAt().Time
		now := time.Now()
		years := int64(now.Sub(created).Hours() / 24 / 365)

		url := fmt.Sprintf("https://img.shields.io/badge/Years-%d-%s?style=%s&logo=%s", years, color, style, logo)

		svg := getSVG(url)

		w.Header().Set("content-type", "image/svg+xml")
		w.Header().Set("cache-control", "no-cache")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", svg)
	}
}

// Badge with the number of public repos you have
func reposHandler(client github.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos" {
			http.Error(w, "404 not found", http.StatusNotFound)
		}
		if r.Method != "GET" {
			http.Error(w, "Method is not supported", http.StatusNotFound)
		}

		values := r.URL.Query()
		username := values.Get("username")

		color := values.Get("color")
		if len(color) == 0 {
			color = "brigthgreen"
		}

		style := values.Get("style")

		logo := values.Get("logo")

		// Fetch the user data
		// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user
		user, _, err := client.Users.Get(context.Background(), username)
		if err != nil {
			log.Fatal("Error fetching user\n %w", username, err)
		}

		// Get the number of public repos the user has
		repos := user.GetPublicRepos()

		url := fmt.Sprintf("https://img.shields.io/badge/Repos-%d-%s?style=%s&logo=%s", repos, color, style, logo)

		svg := getSVG(url)

		w.Header().Set("content-type", "image/svg+xml")
		w.Header().Set("cache-control", "no-cache")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", svg)

	}
}

// Get SVG from shields.io for rendering
func getSVG(url string) []byte {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal("Error fetching badge from shields.io")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading the GET body")
	}

	return data
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
	http.HandleFunc("/organizations", organizationsHandler(*client))

	// Show number of years user has been a GitHub member
	http.HandleFunc("/years", yearsHandler(*client))

	http.HandleFunc("/repos", reposHandler(*client))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server\n %w", err)
	}
}
