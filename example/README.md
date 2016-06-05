## Using this Example

[This](example/example.go), a simple HTTP server which creates the reCaptcha form and tests the input.

Build the example after installing the recaptcha package:

```
go get github.com/dpapathanasiou/go-recaptcha
go build example.go
```

Run the server by invoking the executable:

```
./example <reCaptcha public key> <reCaptcha private key>
```

You can access the page from http://localhost:9001/ in your browser.