package main

import (
	"log"

	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
)

// GetArchmagesList returns the data from the archmage leaderboard
func (a *App) ArchmagesList() []bootdevapi.Archmage {
	list, err := bootdevapi.Archmages(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

// GlobalStats returns the general global stats from the leaderboard
func (a *App) GlobalStats() bootdevapi.GlobalStats {
	stats, err := bootdevapi.GetGeneralStats(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return stats
}

// TopDailyLearners returns the top 30 users based on exp earned
func (a *App) TopDailyLearners() []bootdevapi.LeaderboardUser {
	list, err := bootdevapi.GetDailyStats(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

// TopCommunity returns the top 30 members of the discord community,
// based on a variety of factors such as activity
func (a *App) TopCommunity() []bootdevapi.Archon {
	list, err := bootdevapi.GetDiscordLeaderboard(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}
