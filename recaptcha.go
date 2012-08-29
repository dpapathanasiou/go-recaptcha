// Package recaptcha handles reCaptcha (http://www.google.com/recaptcha) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers reCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the recaptcha_private_key constant before building and using
package recaptcha

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const recaptcha_server_name = "http://www.google.com/recaptcha/api/verify"

var recaptcha_private_key string

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func check(remoteip, challenge, response string) (s string) {
	s = ""
	resp, err := http.PostForm(recaptcha_server_name,
		url.Values{"privatekey": {recaptcha_private_key}, "remoteip": {remoteip}, "challenge": {challenge}, "response": {response}})
	if err != nil {
		log.Println("Post error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body: %s", err)
	} else {
		s = string(body)
	}
	return
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func Confirm(remoteip, challenge, response string) (result bool) {
	result = strings.HasPrefix(check(remoteip, challenge, response), "true")
	return
}

// Init allows the webserver or code evaluating the reCaptcha form input to set the
// reCaptcha private key (string) value, which will be different for every domain.
func Init(key string) {
	recaptcha_private_key = key
}
