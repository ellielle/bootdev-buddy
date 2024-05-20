package main

import (
	"context"
	"log"
	"time"

	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
	"github.com/ellielle/bootdev-buddy/internal/cache"
)

// App struct
type App struct {
	ctx   context.Context
	cache cache.Cache
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.cache = cache.NewCache(5 * time.Minute)
}

// GetArchmagesList returns the data from the archmage leaderboard
func (a *App) ArchmagesList() []bootdevapi.Archmage {
	list, err := bootdevapi.GetArchmages(a.cache)
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

func (a *App) TopDailyLearners() []bootdevapi.LeaderboardUser {
	list, err := bootdevapi.GetDailyStats(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func (a *App) TopCommunity() []bootdevapi.Archsage {
	list, err := bootdevapi.GetDiscordLeaderboard(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}
