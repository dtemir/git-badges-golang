package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func getGitHubClient(token string) *github.Client {
	// Create a GitHub client to make API calls
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)

	return ghClient
}

func getMongoDBCollection() *mongo.Collection {
	// Create a MongoDB client to store visit counts
	mgClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal("MongoDB: Couldn't create a MongoDB client\n %w", err)
	}

	// Ping to test if the MongoDB database has been found
	if err := mgClient.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("MongoDB: Couldn't find a MongoDB database\n %w", err)
	}

	visitsCollection := mgClient.Database("production").Collection("visits")

	return visitsCollection
}

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

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server\n %w", err)
	}
}
