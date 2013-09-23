	Under heavy construction. Pet project to learn Go

# Goryproxy

It aims at simulating failures on services by proxying to various services.

Example of failures that will be considered:

* tcp: disconnection, connection hanging
* http: intermittent error 500
* mysql: connection lost, too many connections open

### Installation

Install packages:

	$ go get github.com/bitly/go-simplejson

### Run

	$ go build
	$ ./goryproxy configs/http.json
