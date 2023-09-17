package algorithm

import (
	"context"
	"fmt"

	"github.com/google/go-github/v55/github"
)

func Bfs(ctx context.Context, client *github.Client, userOrigin, userTarget string, maxLevel, maxFollowersToVisitPerUser int) ([]string, int, error) {
	queue := []string{userOrigin}
	visited := make(map[string]bool)
	pathMap := make(map[string][]string)
	level := 0

	if userOrigin == userTarget {
		return []string{userOrigin}, 0, nil
	}

	for level <= maxLevel && len(queue) > 0 {
		nextQueue := []string{}

		for _, currentUser := range queue {
			fmt.Printf("Visiting user: %s | Current level: L%d\n", currentUser, level)
			visited[currentUser] = true

			followers, _, err := client.Users.ListFollowers(ctx, currentUser, &github.ListOptions{PerPage: maxFollowersToVisitPerUser})
			if err != nil {
				return nil, 0, err
			}

			for _, follower := range followers {
				followerUser := follower.GetLogin()

				if !visited[followerUser] {
					nextQueue = append(nextQueue, followerUser)

					pathMap[followerUser] = append(pathMap[currentUser], followerUser)

					if followerUser == userTarget {
						return pathMap[followerUser], level, nil
					}
				}
			}
		}

		queue = nextQueue
		level++
	}

	return nil, 1, fmt.Errorf("No path found between %s and %s within %d degrees", userOrigin, userTarget, maxLevel)
}
