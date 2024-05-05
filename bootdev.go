package main

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getBDLeaderboard() {
	const LEADERBOARD = "https://boot.dev/leaderboard"

	req, err := http.NewRequest("GET", LEADERBOARD, nil)
	if err != nil {
		log.Fatalf("Couldn't create request %v", err.Error())
	}

	// TODO: find the href "/u/"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Couldn't get leaderboard: %v", err.Error())
	}

	body, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing html: %v", err.Error())
	}

	// matching function to only grab data from the Archmage section
	matcher := func(node *html.Node) (keep bool, exit bool) {
		if node.Type == html.ElementNode && strings.TrimSpace(node.Data) != "" {
			keep = true
		}
		if node.DataAtom == atom.Svg {
			exit = true
		}
		return
	}

	nodes := traverseNodes(body, matcher)
	for i, node := range nodes {

	}

}
