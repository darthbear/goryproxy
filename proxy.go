package main

type Proxy interface {
	Listen(proxyHost string, proxyPort int, listenHost string, listenPort int, opts map[string]interface{}) (string, error)
}
