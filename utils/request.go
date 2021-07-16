package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, errors.New("Could not complete the request")
	}

	return res.Body, nil
}

func MakeRequestCon(url string, ch chan<-[]byte) {
	res, err := http.Get(url)
	
	if (err != nil) {
		log.Fatal(err)
	}
	
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	ch <- body
}
