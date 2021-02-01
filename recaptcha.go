package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type recaptchaResponse struct {
	Success    bool      `json:"success"`
	Score      float32   `json:"score"`
	Action     string    `json:"action"`
	Challenge  time.Time `json:"challenge_ts"`
	Hostname   string    `json:"hostname"`
	ErrorCodes []string  `json:"error-codes"`
}

const (
	requestTimeout  = time.Second * 10
	recaptchaServer = "https://www.google.com/recaptcha/api/siteverify"
	responseKey = "g-recaptcha-response"
)

type Recaptcha struct {
	PrivateKey string
}

func (r *Recaptcha) Verify(token string, action string, score float32) (bool, error) {
	if score < 0 || score > 1 {
		return false, errors.New("score must be a number between 0.0 and 1.0")
	}

	httpClient := http.Client{
		Timeout: requestTimeout,
	}

	resp, err := httpClient.PostForm(
		recaptchaServer,

		url.Values{
			"secret": {r.PrivateKey},
			"response": {token},
		},
	)

	defer resp.Body.Close()
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var captcha recaptchaResponse
	err = json.Unmarshal(body, &captcha)
	if err != nil {
		return false, err
	}

	if strings.ToLower(captcha.Action) != strings.ToLower(action) {
		return false, errors.New("reCAPTCHA actions do not match")
	}

	if captcha.Score < score {
		return false, nil
	}

	return captcha.Success, nil
}


