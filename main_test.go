package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Test index redirecting to the repo for instructions on use
func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	indexHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	// Status must be 301 Moved Permanently to redirect to the repo
	if res.Status != "301 Moved Permanently" {
		t.Errorf("Test: Unsuccessful status code | Expected %s Received %s", "301 Moved Permanently", res.Status)
	}

	// Location must point to the repo name
	redirectURL := "https://github.com/dtemir/git-badges-golang"
	if res.Header.Get("Location") != redirectURL {
		t.Errorf("Test: Location isn't setup correctly | Expected %s Received | %s", redirectURL, res.Header.Get("Location"))
	}
}

// Test organization endpoint to return 200 and SVG
func TestOrganizationsHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n %w", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/organizations?username=dtemir", nil)
	w := httptest.NewRecorder()

	ghClient := getGitHubClient(os.Getenv("GITHUB_TOKEN"))
	organizationsHandler(*ghClient)(w, req)

	res := w.Result()
	defer res.Body.Close()

	// Status must be 200 OK
	if res.Status != "200 OK" {
		t.Errorf("Test: Unsuccessful status code | Expected %s Received %s", "200 OK", res.Status)
	}

	// Body must contain two matching SVG tags
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Test: Expected nil Received %v", err)
	}
	count := getSVGCount(string(data))
	if count != 1 {
		t.Errorf("Test: SVG tag count incorrect | Expected %d Received %d", 1, count)
	}
}

// Test years endpoint to return 200 and SVG
func TestYearsHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n %w", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/years?username=dtemir", nil)
	w := httptest.NewRecorder()

	ghClient := getGitHubClient(os.Getenv("GITHUB_TOKEN"))
	yearsHandler(*ghClient)(w, req)

	res := w.Result()
	defer res.Body.Close()

	// Status must be 200 OK
	if res.Status != "200 OK" {
		t.Errorf("Test: Unsuccessful status code | Expected %s Received %s", "200 OK", res.Status)
	}

	// Body must contain two matching SVG tags
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Test: Expected nil Received %v", err)
	}
	count := getSVGCount(string(data))
	if count != 1 {
		t.Errorf("Test: SVG tag count incorrect | Expected %d Received %d", 1, count)
	}
}

// Test repos endpoint to return 200 and SVG
func TestReposHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n %w", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/repos?username=dtemir", nil)
	w := httptest.NewRecorder()

	ghClient := getGitHubClient(os.Getenv("GITHUB_TOKEN"))
	reposHandler(*ghClient)(w, req)

	res := w.Result()
	defer res.Body.Close()

	// Status must be 200 OK
	if res.Status != "200 OK" {
		t.Errorf("Test: Unsuccessful status code | Expected %s Received %s", "200 OK", res.Status)
	}

	// Body must contain two matching SVG tags
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Test: Expected nil Received %v", err)
	}
	count := getSVGCount(string(data))
	if count != 1 {
		t.Errorf("Test: SVG tag count incorrect | Expected %d Received %d", 1, count)
	}
}
