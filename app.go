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
func (a *App) GetArchmagesList() []bootdevapi.Archmage {
	list, err := bootdevapi.GetArchmages(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

// GetStats returns the general global stats from the leaderboard
func (a *App) GetStats() bootdevapi.GlobalStats {
	stats, err := bootdevapi.GetGeneralStats(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return stats
}
