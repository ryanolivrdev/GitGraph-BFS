package main

import (
	"backend/algorithm"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

const accessToken = ""

var (
	maxLevel                   = 1000
	maxFollowersToVisitPerUser = 1000
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userOrigin := r.URL.Query().Get("userOrigin")
		userTarget := r.URL.Query().Get("userTarget")

		path, kevinBaconNumber, err := algorithm.Bfs(ctx, client, userOrigin, userTarget, maxLevel, maxFollowersToVisitPerUser)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Println(path, kevinBaconNumber)
		fmt.Fprintf(w, "{\"path\": [%v], \"kevinBaconNumber\": %d}", path, kevinBaconNumber)
	})

	fmt.Println("listening on", port)
	http.ListenAndServe(":"+port, nil)
}
