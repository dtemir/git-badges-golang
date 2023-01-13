package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/google/go-github/v48/github"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/oauth2"
)

// Get a GitHub client to make API calls
func getGitHubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)

	return ghClient
}

// Get a MongoDB client to store visit counts
func getMongoDBCollection() *mongo.Collection {
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

// Get the visit count from MongoDB given the repo name
func getVisitsCount(repo string, collection mongo.Collection) int64 {
	var visits int64

	// Find the document to update
	var err error
	var doc bson.M
	filter := bson.M{"name": repo}
	err = collection.FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		// No document exist, i.e. first time visit
		if err == mongo.ErrNoDocuments {
			visits = 1
			// Create a new document with one visit (make sure it is int64)
			doc = bson.M{"name": repo, "visits": int64(visits)}
			_, err = collection.InsertOne(context.Background(), doc)
			if err != nil {
				log.Fatal("MongoDB: Failed to very first visit\n %w", err)
			}
		} else {
			log.Fatal("MongoDB: Couldn't find the document to update\n %w", err)
		}
	} else {
		// Increment the number of visits by one
		v, ok := doc["visits"].(int64)
		if !ok {
			log.Fatal("MongoDB: Error converting 'visits' field to int64\n %w", err)
		}
		doc["visits"] = v + 1

		visits = v + 1

		// Update the document in the collection with a new count
		_, err = collection.ReplaceOne(context.Background(), filter, doc)
		if err != nil {
			log.Fatal("MongoDB: Couldn't update the document in the collection")
		}
	}

	return visits
}

// Get number of SVG tags for testing
// Example: <svg...>...</svg>
func getSVGCount(svg string) int {
	re := regexp.MustCompile(`<svg(.*?)</svg>`)
	matches := len(re.FindAllString(svg, -1))

	return matches
}
