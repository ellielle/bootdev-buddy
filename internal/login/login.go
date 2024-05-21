package login

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ExchangeOTPForToken(OTP string) (string, error) {
	// Boot.Dev's one-time password login for their CLI app
	// The token returned can be used to authenticate the user
	// across their API
	const OTP_LOGIN = "https://api.boot.dev/v1/auth/otp/login"

	type apiRequest struct {
		Otp string `json:"Otp"`
	}

	type loginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	// populate the struct and then Marshal it into a []byte
	// to be used in the HTTP request
	apiReq := apiRequest{Otp: OTP}
	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		return "", err
	}

	// Create and carry out a new HTTP request to the API
	req, err := http.NewRequest("POST", OTP_LOGIN, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	loginResp := loginResponse{}
	err = json.Unmarshal(token, &loginResp)
	if err != nil {
		return "", err
	}

	return loginResp.AccessToken, nil
}
