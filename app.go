package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
	"github.com/ellielle/bootdev-buddy/internal/cache"
	"github.com/ellielle/bootdev-buddy/internal/login"
)

// App struct
type App struct {
	ctx   context.Context
	cache cache.Cache
	token string
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
func (a *App) TopCommunity() []bootdevapi.Archsage {
	list, err := bootdevapi.GetDiscordLeaderboard(a.cache)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func (a *App) LoginUser(OTP string) (bool, error) {
	token, err := login.ExchangeOTPForToken(OTP)
	if err != nil {
		return false, errors.New("error exchanging OTP for Token")
	}
	if token.AccessToken == "" {
		return false, errors.New("empty token after exchanging with OTP")
	}

	// set token in App struct so it can be used for
	// user-specific queries
	a.token = token.AccessToken

	return true, nil
}

func (a *App) UserData() {
	if a.token == "" {
		return
	}

}
