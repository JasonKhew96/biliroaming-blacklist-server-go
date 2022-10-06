package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (w *Web) verifyCaptchas(token, ip string) (bool, error) {
	reqUrl := "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	body := url.Values{
		"remoteip": {ip},
		"response": {token},
		"secret":   {w.config.SecretKey},
	}
	req, err := http.NewRequest(http.MethodPost, reqUrl, strings.NewReader(body.Encode()))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return false, fmt.Errorf("content type: %s", resp.Header.Get("Content-Type"))
	}
	var data struct {
		Success     bool      `json:"success"`
		// ChallengeTs time.Time `json:"challenge_ts"`
		// Hostname    string    `json:"hostname"`
		// ErrorCodes  []string  `json:"error-codes"`
		// Action      string    `json:"action"`
		// CData       string    `json:"cdata"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return false, err
	}
	return data.Success, nil
}
