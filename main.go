package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	// Setup logging
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Setup: Couldn't create logs.txt file\n %v", err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Load in .env file with GitHub Token
	err := godotenv.Load()
	if err != nil {
		ErrorLogger.Fatal("Setup: Error loading .env file\n %w", err)
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
		ErrorLogger.Fatal("Setup: Error starting the server\n %w", err)
	}
}
