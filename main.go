package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var asn string
	fmt.Scanln(&asn)
	url := fmt.Sprintf("https://peeringdb.com/api/net?asn=%s", asn)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	network, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", network)
}
