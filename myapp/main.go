package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/v36/github"
	"golang.org/x/oauth2"
)

const (
	clientID     = "Iv1.3148739c0ce7a95f"
	clientSecret = "ef1a3e883aaa4de32a143bbf4aaae821ca8b42a1"
	redirectURI  = "http://localhost:3001/oauth-callback"
	appID        = 716142
)

var oauthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  redirectURI,
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	},
}

type RepoData struct {
	Name string
}

var username string

func main() {

    

	http.HandleFunc("/oauth-callback", handleOAuthCallback)
    http.HandleFunc("/get-username", handleGetUsername)

	port := 4001
	log.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}



func handleGetUsername(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"username": "%s"}`, username)
}


// Modify handleOAuthCallback function
func handleOAuthCallback(w http.ResponseWriter, r *http.Request)  {
	code := r.URL.Query().Get("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return 
	}

	// Use the token to get user information from GitHub
	client := github.NewClient(oauthConfig.Client(context.Background(), token))
    user, _, err := client.Users.Get(context.Background(), "")
	repos, _, err := client.Repositories.List(context.Background(), "", nil)
    username = *user.Login
    fmt.Println(username)
	if err != nil {
		http.Error(w, "Failed to get repositories from GitHub", http.StatusInternalServerError)
		return 
	}

	// Extract repository names
	var repoData []RepoData
	for _, repo := range repos {
		repoData = append(repoData, RepoData{Name: *repo.Name})
	}

	// Read the content of the external HTML file
	htmlContent, err := ioutil.ReadFile("repos.html")
	if err != nil {
		http.Error(w, "Failed to read HTML file: "+err.Error(), http.StatusInternalServerError)
		return 
	}

	tmpl, err := template.New("repos").Parse(string(htmlContent))
	if err != nil {
		http.Error(w, "Failed to parse HTML template: "+err.Error(), http.StatusInternalServerError)
		return 
	}

	data := struct {
		Repos []RepoData
	}{
		Repos: repoData,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute HTML template: "+err.Error(), http.StatusInternalServerError)
		return 
	}
    
}


