package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var network_id string
	fmt.Scanln(&network_id)
	url := fmt.Sprintf("https://www.peeringdb.com/api/net/%s", network_id)
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
