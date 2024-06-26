package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type BDToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const LOCAL_KEYS = ".bootdevbuddy.json"

const REFRESH_URL = "https://api.boot.dev/v1/auth/refresh"

const OTP_LOGIN = "https://api.boot.dev/v1/auth/otp/login"

func ExchangeOTPForToken(OTP string) (*BDToken, error) {
	if OTP == "" {
		return nil, errors.New("empty one-time password")
	}
	// Boot.Dev's one-time password login for their CLI app
	// The token returned can be used to authenticate the user
	// across their API

	type apiRequest struct {
		Otp string `json:"Otp"`
	}

	// populate the struct and then Marshal it into a []byte
	// to be used in the HTTP request
	apiReq := apiRequest{Otp: OTP}
	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}

	// Create and carry out a new HTTP request to the API
	req, err := http.NewRequest("POST", OTP_LOGIN, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	loginResp := BDToken{}
	err = json.Unmarshal(token, &loginResp)
	if err != nil {
		return nil, err
	}

	return &loginResp, err
}

// RefreshToken gets the refresh_token from the local file
// and uses it to refresh the user's access_token.
func RefreshToken() (*BDToken, error) {
	tokens, err := readKeys()
	if err != nil {
		return nil, err
	}

	tokenData, err := json.Marshal(tokens)
	if err != nil {
		return nil, err
	}

	var jsonData BDToken
	err = json.Unmarshal(tokenData, &jsonData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", REFRESH_URL, bytes.NewBuffer(tokenData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Refresh-Token", jsonData.RefreshToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("invalid refresh token")
	}

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newTokens BDToken
	err = json.Unmarshal(token, &newTokens)
	if err != nil {
		return nil, err
	}

	return &newTokens, err
}

func SaveTokens(tokens *BDToken) error {
	err := writeKeys(tokens)

	return err
}

func ReadTokens() (*BDToken, error) {
	tokens, err := readKeys()

	return tokens, err
}

func LogoutUser() error {
	err := deleteKeys()

	return err
}

// deleteKeys deletes the file, preventing logging in
// automatically on launch
func deleteKeys() error {
	_, err := os.Stat(LOCAL_KEYS)

	if os.IsNotExist(err) {
		return nil
	}

	err = os.Remove(LOCAL_KEYS)
	return err
}

// writeKeys takes the Boot.Dev access_token and
// refresh_token and saves them locally.
func writeKeys(tokens *BDToken) error {
	_, err := os.Stat(LOCAL_KEYS)

	if os.IsNotExist(err) {
		_, err = os.Create(LOCAL_KEYS)
		if err != nil {
			return err
		}
	}

	keys, err := json.Marshal(&tokens)
	if err != nil {
		return errors.New("error marshaling json")
	}

	err = os.WriteFile(LOCAL_KEYS, keys, 0600)

	return err
}

// readKeys reads .bootdevbuddy.json, unmarshals it
// into a struct, and returns the access_token and
// refresh_token
func readKeys() (*BDToken, error) {
	_, err := os.Stat(LOCAL_KEYS)

	if os.IsNotExist(err) {
		return nil, errors.New("no saved data")
	}

	file, err := os.ReadFile(LOCAL_KEYS)
	if err != nil {
		return nil, errors.New("error opening key file")
	}

	var keys BDToken
	err = json.Unmarshal(file, &keys)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}
