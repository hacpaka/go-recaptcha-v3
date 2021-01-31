sgo-recaptcha-v3
============

About
-----

This package is designed to handle GOOGLE [reCAPTCHA v3](https://developers.google.com/recaptcha/intro) in [Go](http://golang.org/).

Usage
-----

Install the package in your environment:

```shell
go get github.com/hacpaka/go-recaptcha-v3
```

Don't forget to import it:
```go
import (
    "github.com/hacpaka/go-recaptcha-v3"
)
```

Somewhere in you code: 
```go
recaptcha := Recaptcha{ PrivateKey: "your-recaptcha-private-key" }
result, err := recaptcha.Verify(action string, response string, minScore uint)
```
