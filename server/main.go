package main

import (
	"backend/algorithm"
	"context"
	"fmt"
	"net/http"
	"os"
	"encoding/json"

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

		if userOrigin == "" || userTarget == "" {
			fmt.Printf("Error: %v\n", "Parametros invalidos")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			data := map[string]interface{}{
				"error": "Parametros invalidos",
			}
			json.NewEncoder(w).Encode(data)
			return
		}

		fmt.Println(userOrigin, userTarget)
		//NORMALIZAR
		

		_, _, err := client.Users.Get(ctx, userOrigin)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			data := map[string]interface{}{
				"error": "Usuario de origen nao encontrado",
			}
			json.NewEncoder(w).Encode(data)
			return
		}

		_, _, err = client.Users.Get(ctx, userTarget)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			data := map[string]interface{}{
				"error": "Usuario de destino nao encontrado",
			}
			json.NewEncoder(w).Encode(data)
			return
		}

		path, kevinBaconNumber, err := algorithm.Bfs(ctx, client, userOrigin, userTarget, maxLevel, maxFollowersToVisitPerUser)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			data := map[string]interface{}{
				"error": err,
			}
			json.NewEncoder(w).Encode(data)
			return
		}

		fmt.Println(path, kevinBaconNumber)
		w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"path":               path,
			"kevinBaconNumber":   kevinBaconNumber,
			"maxLevel":           maxLevel,
			"maxFollowersToVisitPerUser": maxFollowersToVisitPerUser,
		}
    json.NewEncoder(w).Encode(data)
	})

	fmt.Println("listening on", port)
	http.ListenAndServe(":"+port, nil)
}
