go-recaptcha
============

About
-----

This package handles reCaptcha (http://www.google.com/recaptcha) form submissions in Go (http://golang.org/).

Usage
-----

Install Clone this repo, add its path to your $GOPATH environment variable, and edit the recaptcha_private_key constant in line 19 of the src/recaptcha/recaptcha.go file to the one provided for your domain.

Install the package in your environment:

```
go get github.com/dpapathanasiou/go-recaptcha
```

To use it within your own code, import "github.com/dpapathanasiou/go-recaptcha" and call:

```
recaptcha.Init (recaptcha_private_key)
```

once, to set the reCaptcha private key for your domain, then:

```
recaptcha.Confirm (client_ip_address, recaptcha_challenge_field, recaptcha_response_field)
```

for each reCpatcha form input you need to check, using the values obtained by reading the form's POST parameters.

The recaptcha.Confirm() function returns either true (i.e., the captcha was completed correctly) or false.

Usage Example
-------------

Included with this repo is example.go, a simple HTTP server which creates the reCaptcha form and tests the input.

Build the example after installing the recaptcha package:

```
go build example.go
```

Run the server by invoking the executable:

```
./example <reCaptcha public key> <reCaptcha private key>
```

You can access the page from http://localhost:9001/ in your browser.

