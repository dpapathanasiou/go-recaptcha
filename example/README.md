## Using this Example

[This](example.go) is a simple HTTP server which creates the reCaptcha form and tests the input.

Build the example after installing the recaptcha package:

```
go get github.com/dpapathanasiou/go-recaptcha
cd $GOPATH/src/github.com/dpapathanasiou/go-recaptcha/example
go build example.go
```

Run the server<sup>&#42;</sup> by invoking the executable:

```
./example <reCaptcha public key (aka Site key)> <reCaptcha private key (aka Secret key)>
```

You can access the page from http://localhost:9001/ in your browser.

For more information on client side setup and other configuration options, check the [official documentation](https://developers.google.com/recaptcha/intro).

 <sup>&#42;</sup> make sure ['localhost' is added to the list of domains allowed](https://developers.google.com/recaptcha/docs/domain_validation) for the site registered at reCaptcha.
