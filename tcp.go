package main

import (
	"log"
)

type TcpProxy struct {
}

func (t TcpProxy) Listen(proxyHost string, proxyPort int, listenHost string, listenPort int, options map[string]interface{}) (string, error) {
	log.Println("Running TCP proxy...")
	return "done", nil
}
