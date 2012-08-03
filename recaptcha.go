
package recaptcha

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
    "log"
)

const (
    recaptcha_server_name = `http://www.google.com/recaptcha/api/verify`
    recaptcha_private_key = `...[your private key goes here]...`
)

// Return the recaptcha reply string for this client's challenge responses
func check (remoteip, challenge, response string) (s string) {
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

// Return true/false based on the recaptcha server's reply
func Confirm (remoteip, challenge, response string) (result bool) {
    result = strings.HasPrefix(check(remoteip, challenge, response), "true") 
    return
}

