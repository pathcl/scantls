package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// this should test weak ciphers
// TLS_RSA_WITH_3DES_EDE_CBC_SHA is the cipher we want to test
func doRequest(url string, cipher []uint16) {

	// how will I check every cipher ?
	// create a transport for http.client in order to pass tls/options
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			CipherSuites:       cipher,
		},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)

	if err != nil {
		log.Println(url, "is OK")
		return
	} else {
		log.Printf("Check your ciphers! %d should not be enabled", cipher)
	}

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

}

func main() {
	// we need to find a way to test every cipher not only this one
	doRequest(os.Args[1], []uint16{tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA})

}
