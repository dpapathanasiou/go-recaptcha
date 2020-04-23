go-recaptcha
============

https://godoc.org/github.com/dpapathanasiou/go-recaptcha

About
-----

This package handles [reCaptcha](https://www.google.com/recaptcha) (API versions [2](https://developers.google.com/recaptcha/intro) and [3](https://developers.google.com/recaptcha/docs/v3)) form submissions in [Go](http://golang.org/).

Usage
-----

Install the package in your environment:

```
go get github.com/dpapathanasiou/go-recaptcha
```

To use it within your own code, import <tt>github.com/dpapathanasiou/go-recaptcha</tt> and call:

```
recaptcha.Init (recaptchaPrivateKey)
```

once, to set the reCaptcha private key for your domain, then:

```
recaptcha.Confirm (clientIpAddress, recaptchaResponse)
```

### [reCAPTCHA v2](https://developers.google.com/recaptcha/intro)
For each reCaptcha form input you need to check, using the values obtained by reading the form's POST parameters (the <tt>recaptchaResponse</tt> in the above corresponds to the value of <tt>g-recaptcha-response</tt> sent by the reCaptcha server.)

The recaptcha.Confirm() function returns either true (i.e., the captcha was completed correctly) or false, along with any errors (from the HTTP io read or the attempt to unmarshal the JSON reply).

### [reCAPTCHA v3](https://developers.google.com/recaptcha/docs/v3)

Version 3 works differently: instead of interrupting page visitors with a prompt, it runs in the background, computing a score.

This repo has been updated to handle the [score and action in the response](recaptcha.go#L20), but the usage example is still in terms of version 2.

Usage Example
-------------

Included with this repo is [example.go](example/example.go), a simple HTTP server which creates the reCaptcha form and tests the input.

See the [instructions](example/README.md) for running the example for more details.
