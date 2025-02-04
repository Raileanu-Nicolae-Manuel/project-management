package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v68/github"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rootFile := os.Getenv("ROOT_FILE")

	fmt.Println("Root File: ", rootFile)
	if _, err := os.Stat(rootFile); os.IsNotExist(err) {
		fmt.Println(rootFile, "does not exist")
	} else {
		fmt.Println("The provided directory named", rootFile, "exists")
	}

	fileKey := os.Getenv("PRIVATE_KEY")
	if fileKey == "" {
		panic("File Key is empty")
	}

	if os.Getenv("APP_ID") == "" {
		panic("App ID is empty")
	}

	appId, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		panic("App ID is not a valid integer")
	}
	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		panic("Installation ID is not a valid integer")
	}
	// file exists

	// TODO: check if project are in the file

	// TODO: get the repository form github

	// =======================================================================================================================================
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appId, installationID, fileKey)

	// Or for endpoints that require JWT authentication
	// itr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, appId, fileKey)

	if err != nil {
		// Handle error.
		panic(err)
	}

	// Use installation transport with client.
	client := github.NewClient(&http.Client{Transport: itr})

	repository, response, err := client.Repositories.ListAll(context.Background(), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Print("Response Status Code:", response.StatusCode)
	fmt.Print("Response Body:", response)
	fmt.Println("Repositories: ", repository)
	for _, repo := range repository {
		fmt.Println(*(repo.Name))
	}

	// client.

	// =======================================================================================================================================

	// TODO: check the user, will ned multiple staff here

	// TODO: clone the projects from github

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		event := r.Header.Get("X-GitHub-Event")
		fmt.Println("Event: ", event)
		fmt.Println(r.Body)
		json.NewEncoder(w).Encode("ok")
	})
	log.Fatal(http.ListenAndServe(":8765", nil))

}
