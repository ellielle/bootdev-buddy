package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
)

const SALMON_URL = "https://api.boot.dev/v1/user_baked_salmon"

const STONES_URL = "https://api.boot.dev/v1/user_seer_stones"

const XP_URL = "https://api.boot.dev/v1/user_xp_potion"

const FLAME_URL = "https://api.boot.dev/v1/user_frozen_flame"

// SeerStones gets the user's current amount of Seer Stones
func (a *App) SeerStones() (int, error) {
	type repsonse struct {
		Stones int `json:"stones"`
	}

	req, err := http.NewRequest("GET", STONES_URL, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.tokens.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	stones, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(string(stones))

	return num, err

}

func (a *App) BakedSalmon() (int, error) {
	type response struct {
		Salmon int `json:"salmon"`
	}

	req, err := http.NewRequest("GET", SALMON_URL, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.tokens.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	salmon, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(string(salmon))
	return num, err
}

func (a *App) XPPotion() (bootdevapi.PotionList, error) {
	var potions bootdevapi.PotionList

	req, err := http.NewRequest("GET", XP_URL, nil)
	if err != nil {
		return bootdevapi.PotionList{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.tokens.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return bootdevapi.PotionList{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&potions)
	if err != nil {
		return bootdevapi.PotionList{}, err
	}

	return potions, nil
}

func (a *App) FrozenFlames() (int, error) {
	req, err := http.NewRequest("GET", FLAME_URL, nil)
	if err != nil {
		return 0, err
	}

	var frozenFlames []bootdevapi.FrozenFlame

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.tokens.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	log.Printf("frozen flames: %#v", resp.Body)

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&frozenFlames)

	// only need the length of the FrozenFlames slice
	return len(frozenFlames), nil
}
