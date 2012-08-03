go-recaptcha
============

About
-----

This package handles reCaptcha (http://www.google.com/recaptcha) form submissions in Go (http://golang.org/).

Usage
-----

Import "github.com/dpapathanasiou/go-recaptcha" in your web server, set the recaptcha_private_key variable to the one provided for your domain, and call:

    recaptcha.Confirm (client_ip_address, recaptcha_challenge_field, recaptcha_response_field)

with the values obtained by reading the form's POST parameters.

The recaptcha.Confirm() function returns either true (i.e., the captcha was completed correctly) or false.

Usage Example
-------------

This is a simple HTTP server which creates the recaptcha form and tests the input.

Set the recaptcha_public_key constant to your actual public key, and after building, run ./example from a prompt. 

You can access the page from http://localhost:9001/ in your browser.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/dpapathanasiou/go-recaptcha"
)

const (
    recaptcha_public_key = `...[your public key goes here]...`
    recaptcha_server_form = `https://www.google.com/recaptcha/api/challenge`
    pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#ff0000;} .ack{color:#0000ff;}</style></head><title>Recaptcha Test</title>
<body><h3>Recaptcha Test</h3>
<p>This is a test form for the go-recaptcha package</p>`
    form       = `<form action="/" method="POST">
    	<script src="%s?k=%s" type="text/javascript"> </script>
    	<input type="submit" name="button" value="Ok">
</form>`
    pageBottom = `</body></html>`
    anError    = `<p class="error">%s</p>`
    anAck      = `<p class="ack">%s</p>`
)

func processRequest(request *http.Request) (result bool) {
    result = false
    challenge, challenge_found := request.Form["recaptcha_challenge_field"]
    recaptcha_resp, resp_found := request.Form["recaptcha_response_field"]
    if challenge_found && resp_found {
    	result = recaptcha.Confirm ("127.0.0.1", challenge[0], recaptcha_resp[0])
    }
    return 
}

func homePage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm() // Must be called before writing response
    fmt.Fprint(writer, pageTop)
    if err != nil {
        fmt.Fprintf(writer, fmt.Sprintf(anError, err))
    } else {
    	_, button_clicked := request.Form["button"]
    	if button_clicked {
    		if processRequest(request) {
    			fmt.Fprint(writer, fmt.Sprintf(anAck, "Recaptcha was correct!"))
    		} else {
    			fmt.Fprintf(writer, fmt.Sprintf(anError, "Recaptcha was incorrect; try again."))
    		}
    	}
    }
    fmt.Fprint(writer, fmt.Sprintf(form, recaptcha_server_form, recaptcha_public_key))
    fmt.Fprint(writer, pageBottom)
}

func main() {
    http.HandleFunc("/", homePage)
    if err := http.ListenAndServe(":9001", nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}
```
