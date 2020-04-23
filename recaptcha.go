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
	"log"
	"net/http"
	"net/url"
	"time"
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

// Check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a RecaptchaResponse object wich contains user response.
func Check(remoteip, response string) (r RecaptchaResponse, err error) {
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {recaptchaPrivateKey}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		log.Printf("Post error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body: %s", err)
		return
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println("Read error: got invalid JSON: %s", err)
		return
	}
	return
}

// Confirm calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func Confirm(remoteip, response string) (result bool, err error) {
	resp, err := Check(remoteip, response)
	result = resp.Success
	return
}

// Init allows the webserver or code evaluating the reCaptcha form input to set the
// reCaptcha private key (string) value, which will be different for every domain.
func Init(key string) {
	recaptchaPrivateKey = key
}
