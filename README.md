go-recaptcha-v3
============

About
-----

This package is designed to handle [Google reCAPTCHA v3](https://developers.google.com/recaptcha/intro) in [Go](http://golang.org/).

Usage
-----

Install the package in your environment:

```shell
go get github.com/hacpaka/go-recaptcha-v3
```

Don't forget to import:
```go
import (
    "github.com/hacpaka/go-recaptcha-v3"
)
```

Now somewhere in your code: 
```go
recaptcha := Recaptcha{ PrivateKey: "your-recaptcha-private-key" }
result, err := recaptcha.Verify(action string, response string, minScore uint)
```

Origin
------

This package was originally forked from: [hazcod/gorecaptcha](https://github.com/hazcod/gorecaptcha).

License 
-------

This package is released under the [MIT license](https://github.com/hacpaka/go-recaptcha-v3/blob/master/LICENSE).