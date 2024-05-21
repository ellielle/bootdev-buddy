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
	writeToken(token.AccessToken)

	return true, nil
}

func (a *App) UserData() {
	if a.token == "" {
		return
	}

}

// TODO: check token validity
// writeToken writes the access_token and
// refresh_token to a local file.
// access_token is only valid for 1 hour
func writeToken(token string) error {
	_, err := os.Stat(".bootdevbuddy.json")
	var file *os.File

	if os.IsNotExist(err) {
		file, err = os.Create(".bootdevbuddy.json")
		if err != nil {
			return err
		}
	} else {
		file, err = os.Open(".bootdevbuddy.json")
		if err != nil {
			return err
		}
	}
	defer file.Close()

	file.WriteString(token)

	return nil
}

// TODO: check token validity
// readToken reads the user's access_token and
// refresh_token from the local file. Token will
// fail if stale and will need to be refreshed
func readToken() (string, error) {
	_, err := os.Stat(".bootdevbuddy.json")
	var file *os.File

	if os.IsNotExist(err) {
		file, err = os.Create(".bootdevbuddy.json")
		if err != nil {
			return "", err
		}
	} else {
		file, err = os.Open(".bootdevbuddy.json")
		if err != nil {
			return "", err
		}
	}
	defer file.Close()

	token, err := os.ReadFile(".bootdevbuddy.json")
	if err != nil {
		return "", err
	}

	return string(token), nil
}
