# https

Simple HTTP/HTTPS file server written in Go. All configuration options passed on the command line.

To install, set your [GOPATH](https://golang.org/doc/code.html#GOPATH) and then run:

    go get github.com/drichardson/https

You can then run https with:

    $GOPATH/bin/https

To create a TLS certificate and private key, run:

    openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout key.pem -out cert.pem
