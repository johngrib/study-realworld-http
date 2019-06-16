package recipe

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func RequestGet(uriAddress string) {
	resp, err := http.Get(uriAddress)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

func RequestGetWithQuery(uriAddress string) {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get(uriAddress + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}