package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
	"github.com/ellielle/bootdev-buddy/internal/cache"
	"github.com/ellielle/bootdev-buddy/internal/login"
)

// App struct
type App struct {
	ctx    context.Context
	cache  cache.Cache
	tokens login.BDToken
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
	tokens, err := login.RefreshToken()
	if err != nil {
		return
	}

	err = login.SaveTokens(tokens)
	if err != nil {
		log.Print(err)
	}

	a.tokens = *tokens
	runtime.EventsEmit(ctx, "domready")
}

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

// LoginUserWithOTP takes the user's one-time password and exchanges it
// for an access_token and refresh_token from Boot.Dev
func (a *App) LoginUserWithOTP(OTP string) (bool, error) {
	tokens, err := login.ExchangeOTPForToken(OTP)
	if err != nil {
		return false, errors.New("error exchanging OTP for Token")
	}
	if tokens.AccessToken == "" {
		return false, errors.New("empty token after exchanging with OTP")
	}

	err = login.SaveTokens(tokens)
	if err != nil {
		return false, err
	}
	// set token in App struct so it can be used for
	// user-specific queries
	a.tokens = *tokens

	return true, nil
}

func (a *App) LoginUserWithToken() (bool, error) {
	// Don't waste calls and assume the 1 hour token is invalid
	tokens, err := login.RefreshToken()
	if err != nil {
		return false, err
	}

	a.tokens = *tokens

	return true, nil
}

// UserData sends an authenticated request to gather the user's
// data for display in the app.
func (a *App) UserData() {
	log.Print("UserData")
}
