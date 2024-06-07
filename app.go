package main

import (
	"context"
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

// UserData sends an authenticated request to gather the user's
// data for display in the app.
func (a *App) UserData() (bootdevapi.UserData, error) {
	userData, err := bootdevapi.UserInfo(a.cache, a.tokens.AccessToken)
	if err != nil {
		return bootdevapi.UserData{}, nil
	}

	return userData, nil
}

// BossBattle retrieves the current / most recent boss battle data
func (a *App) BossBattle() (bootdevapi.BossBattle, error) {
	bossData, err := bootdevapi.BossBattleStats(a.cache, a.tokens.AccessToken)
	if err != nil {
		return bootdevapi.BossBattle{}, nil
	}

	return bossData, nil
}
