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
	mg_client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB: Couldn't create a MongoDB client\n %w", err)
	}

	// Ping to test if the MongoDB database has been found
	if err := mg_client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("MongoDB: Couldn't find a MongoDB database\n %w", err)
	}

	visitsCollection := mg_client.Database("production").Collection("visits")

	fmt.Println(visitsCollection)

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
