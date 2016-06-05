// example.go
//
// A simple HTTP server which presents a reCaptcha input form and evaulates the result,
// using the github.com/dpapathanasiou/go-recaptcha package.
//
// See the main() function for usage.
package main

import (
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var recaptchaPublicKey string

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#ff0000;} .ack{color:#0000ff;}</style><title>Recaptcha Test</title></head>
<body><div style="width:100%"><div style="width: 50%;margin: 0 auto;">
<h3>Recaptcha Test</h3>
<p>This is a test form for the go-recaptcha package</p>`
	form = `<form action="/" method="POST">
	    <script src="https://www.google.com/recaptcha/api.js"></script>
			<div class="g-recaptcha" data-sitekey="%s"></div>
    	<input type="submit" name="button" value="Ok">
</form>`
	pageBottom = `</div></div></body></html>`
	anError    = `<p class="error">%s</p>`
	anAck      = `<p class="ack">%s</p>`
)

// processRequest accepts the http.Request object, finds the reCaptcha form variables which
// were input and sent by HTTP POST to the server, then calls the recaptcha package's Confirm()
// method, which returns a boolean indicating whether or not the client answered the form correctly.
func processRequest(request *http.Request) (result bool) {
	result = false
	recaptchaResponse, responseFound := request.Form["g-recaptcha-response"]
	if responseFound {
		result = recaptcha.Confirm("127.0.0.1", recaptchaResponse[0])
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
		_, buttonClicked := request.Form["button"]
		if buttonClicked {
			if processRequest(request) {
				fmt.Fprint(writer, fmt.Sprintf(anAck, "Recaptcha was correct!"))
			} else {
				fmt.Fprintf(writer, fmt.Sprintf(anError, "Recaptcha was incorrect; try again."))
			}
		}
	}
	fmt.Fprint(writer, fmt.Sprintf(form, recaptchaPublicKey))
	fmt.Fprint(writer, pageBottom)
}

// main expects two command-line arguments: the reCaptcha public key for producing the HTML form,
// and the reCaptcha private key, to pass to recaptcha.Init() so the recaptcha package can check the input.
// It launches a simple web server on port 9001 which produces the reCaptcha input form and checks the client
// input if the form is posted.
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s <reCaptcha public key> <reCaptcha private key>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {
		recaptchaPublicKey = os.Args[1]
		recaptcha.Init(os.Args[2])

		http.HandleFunc("/", homePage)
		if err := http.ListenAndServe(":9001", nil); err != nil {
			log.Fatal("failed to start server", err)
		}
	}
}
