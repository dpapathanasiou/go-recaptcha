// Package recaptcha handles reCaptcha (http://www.google.com/recaptcha) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers reCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the recaptchaPrivateKey constant before building and using
package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

var recaptchaPrivateKey string

// check will construct the request to the verification API, send it, and return the result.
func check(remoteip, response string) (RecaptchaResponse, error) {
	var r RecaptchaResponse
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {recaptchaPrivateKey}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		return r, errors.Wrap(err, "post error: %s\n")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, errors.Wrap(err, "read error: could not read body")
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, errors.Wrap(err, "read error: got invalid JSON")
	}
	return r, nil
}

// Confirm is the public interface function that validates the reCAPTCHA token.
// It accepts the client ip address and the token returned to the client after completing the challenge.
// It returns a boolean value indicating whether or not the client token is authentic, meaning the challenge
// was answered correctly.
func Confirm(remoteip, response string) (result bool, err error) {
	resp, err := check(remoteip, response)
	result = resp.Success
	return
}

// Init allows the webserver or code evaluating the reCAPTCHA token input to set the
// reCAPTCHA private key (string) value, which will be different for every domain.
func Init(key string) {
	recaptchaPrivateKey = key
}
