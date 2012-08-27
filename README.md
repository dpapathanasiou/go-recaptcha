go-recaptcha
============

About
-----

This package handles reCaptcha (http://www.google.com/recaptcha) form submissions in Go (http://golang.org/).

Usage
-----

Clone this repo, add its path to your $GOPATH environment variable, and edit the recaptcha_private_key constant in line 19 of the src/recaptcha/recaptcha.go file to the one provided for your domain.

Next, install the package in your environment:

```
cd ~/[where you cloned the repo]/go-recaptcha
export GOPATH=$GOPATH:`pwd`
cd $GOPATH/src/recaptcha
go install
```

To use it within your own code, import "recaptcha" and call:

    recaptcha.Confirm (client_ip_address, recaptcha_challenge_field, recaptcha_response_field)

with the values obtained by reading the form's POST parameters.

The recaptcha.Confirm() function returns either true (i.e., the captcha was completed correctly) or false.

Usage Example
-------------

Included with this repo is example.go, a simple HTTP server which creates the reCaptcha form and tests the input.

Set the recaptcha_public_key constant in line 17 to your actual public key, and build:

```
go build example.go
```

Run the server by invoking the executable:

```
./example
```

You can access the page from http://localhost:9001/ in your browser.

