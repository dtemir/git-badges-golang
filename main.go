package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load in .env file with GitHub Token
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n %w", err)
	}

	ghClient := getGitHubClient(os.Getenv("GITHUB_TOKEN"))

	visitsCollection := getMongoDBCollection()

	// Show number of organizations for user
	http.HandleFunc("/organizations", organizationsHandler(*ghClient))

	// Show number of years user has been a GitHub member
	http.HandleFunc("/years", yearsHandler(*ghClient))

	// Show number of repos user owns
	http.HandleFunc("/repos", reposHandler(*ghClient))

	// Show number of visits to user's GitHub page
	http.HandleFunc("/visits", visitsHandler(*ghClient, *visitsCollection))

	// Show general message if people arrive at index
	http.HandleFunc("/", indexHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server\n %w", err)
	}
}
