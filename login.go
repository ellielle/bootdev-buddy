package main

import (
	"errors"

	"github.com/ellielle/bootdev-buddy/internal/login"
)

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

func (a *App) LogoutUser() error {
	err := login.LogoutUser()
	if err != nil {
		return err
	}

	return nil
}
