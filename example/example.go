// example.go
//
// A simple HTTP server which presents a reCaptcha input form and evaulates the result,
// using the github.com/dpapathanasiou/go-recaptcha package
//
// Edit the recaptcha_public_key constant before using
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"recaptcha"
)

var recaptcha_public_key string

const (
	recaptcha_server_form = `https://www.google.com/recaptcha/api/challenge`
	pageTop               = `<!DOCTYPE HTML><html><head>
<style>.error{color:#ff0000;} .ack{color:#0000ff;}</style></head><title>Recaptcha Test</title>
<body><h3>Recaptcha Test</h3>
<p>This is a test form for the go-recaptcha package</p>`
	form = `<form action="/" method="POST">
    	<script src="%s?k=%s" type="text/javascript"> </script>
    	<input type="submit" name="button" value="Ok">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
	anAck      = `<p class="ack">%s</p>`
)

// processRequest accepts the http.Request object, finds the reCaptcha form variables which 
// were input and sent by HTTP POST to the server, then calls the recaptcha package's Confirm()
// method, which returns a boolean indicating whether or not the client answered the form correctly.
func processRequest(request *http.Request) (result bool) {
	result = false
	challenge, challenge_found := request.Form["recaptcha_challenge_field"]
	recaptcha_resp, resp_found := request.Form["recaptcha_response_field"]
	if challenge_found && resp_found {
		result = recaptcha.Confirm("127.0.0.1", challenge[0], recaptcha_resp[0])
	}
	return
}

// homePage is a simple HTTP handler which produces a basic HTML page
// (as defined by the pageTop and pageBottom constants), including 
// an input form with a reCaptcha challenge.
// If the http.Request object indicates the form input has been posted,
// it calls processRequest() and displays a message indicating whether or not
// the reCaptcha form was input correctly.
// Either way, it writes HTML output through the http.ResponseWriter.
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

// main expects two command-line arguments: the reCaptcha public key for producing the HTML form,
// and the reCaptcha private key, to pass to recaptcha.Init() so the recaptcha package can check the input.
// It launches a simple web server on port 9001 which produces the reCaptcha input form and checks the client
// input if the form is posted.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: %s <reCaptcha public key> <reCaptcha private key>", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {
		recaptcha_public_key = os.Args[1]
		recaptcha.Init(os.Args[2])

		http.HandleFunc("/", homePage)
		if err := http.ListenAndServe(":9001", nil); err != nil {
			log.Fatal("failed to start server", err)
		}
	}
}
