package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type BDTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ExchangeOTPForToken(OTP string) (*BDTokens, error) {
	// Boot.Dev's one-time password login for their CLI app
	// The token returned can be used to authenticate the user
	// across their API
	const OTP_LOGIN = "https://api.boot.dev/v1/auth/otp/login"

	type apiRequest struct {
		Otp string `json:"Otp"`
	}

	// populate the struct and then Marshal it into a []byte
	// to be used in the HTTP request
	apiReq := apiRequest{Otp: OTP}
	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		return &BDTokens{}, err
	}

	// Create and carry out a new HTTP request to the API
	req, err := http.NewRequest("POST", OTP_LOGIN, bytes.NewBuffer(jsonData))
	if err != nil {
		return &BDTokens{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &BDTokens{}, err
	}
	defer resp.Body.Close()

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return &BDTokens{}, err
	}

	loginResp := BDTokens{}
	err = json.Unmarshal(token, &loginResp)
	if err != nil {
		return &BDTokens{}, err
	}

	// save the keys locally, so the user doesn't
	// need to sign in via OTP at regular intervals
	writeKeys(&loginResp)

	return &loginResp, nil
}

// writeKeys takes the Boot.Dev access_token and
// refresh_token and saves them locally.
func writeKeys(tokens *BDTokens) error {
	_, err := os.Stat(".bootdevbuddy.json")
	var file *os.File

	if os.IsNotExist(err) {
		file, err = os.Create(".bootdevbuddy.json")
		if err != nil {
			return errors.New("error creating key file")
		}
	} else {
		file, err = os.Open(".bootdevbuddy.json")
		if err != nil {
			return errors.New("error opening key file")
		}
	}
	defer file.Close()

	keys, err := json.Marshal(tokens)
	if err != nil {
		return errors.New("error marshaling json")
	}

	file.Write(keys)

	return nil
}

func readKeys() (*BDTokens, error) {
	_, err := os.Stat(".bootdevbuddy.json")
	var file *os.File

	if os.IsNotExist(err) {
		file, err = os.Create(".bootdevbuddy.json")
		if err != nil {
			return &BDTokens{}, errors.New("error creating key file")
		}
	} else {
		file, err = os.Open(".bootdevbuddy.json")
		if err != nil {
			return &BDTokens{}, errors.New("error opening key file")
		}
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return &BDTokens{}, errors.New("error reading key file")
	}

	var keys BDTokens
	err = json.Unmarshal(data, &keys)
	if err != nil {
		return &BDTokens{}, errors.New("error unmarshaling json")
	}

	return &keys, nil
}
