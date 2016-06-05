// Package recaptcha handles reCaptcha (http://www.google.com/recaptcha) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers reCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the recaptcha_private_key constant before building and using
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
	Success bool `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname string `json:"hostname"`
	ErrorCodes []int `json:"error-codes"`
}

const recaptcha_server_name = "https://www.google.com/recaptcha/api/siteverify"

var recaptcha_private_key string

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func check(remoteip, response string) (r RecaptchaResponse) {
	resp, err := http.PostForm(recaptcha_server_name,
		url.Values{"secret": {recaptcha_private_key}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		log.Println("Post error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body: %s", err)
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println("Read error: got invalid JSON: %s", err)
	}
	return
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func Confirm(remoteip, response string) (result bool) {
	result = check(remoteip, response).Success
	return
}

// Init allows the webserver or code evaluating the reCaptcha form input to set the
// reCaptcha private key (string) value, which will be different for every domain.
func Init(key string) {
	recaptcha_private_key = key
}
