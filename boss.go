package main

import (
	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
)

// BossBattle retrieves the current / most recent boss battle data
func (a *App) BossBattle() (bootdevapi.BossBattle, error) {
	bossData, err := bootdevapi.BossBattleStats(a.cache, a.tokens.AccessToken)

	if err != nil {
		return bootdevapi.BossBattle{}, err
	}

	return bossData, nil
}
