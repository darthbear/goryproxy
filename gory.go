package main

import (
	"fmt"
	"log"
	. "github.com/bitly/go-simplejson"
	"io/ioutil"
	"os"
)

func readConfig(file string) (*Json, error) {
	var err error
	dat, err := ioutil.ReadFile(file)
	if (err != nil) {
		return nil, err
	}

	js, err := NewJson(dat)
	return js, err
}

func startProxy(conf map[string]interface{}) {
	TcpProxy{}.Listen(conf["proxy_host"].(string),
			int(conf["proxy_port"].(float64)),
			conf["listen_host"].(string),
			int(conf["listen_port"].(float64)),
			nil)
}

func startProxies(root *Json) {
	for _, conf := range root.Get("configs").MustArray() {
		startProxy(conf.(map[string]interface{}))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <config file>\n", os.Args[0])
		os.Exit(1)
	}

	conf, err := readConfig(os.Args[1])
	if (err != nil) {
		log.Fatal(err)
	}

	startProxies(conf)
}
