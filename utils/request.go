package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func MakeRequestHeadless(url string) (string, error) {
	l := launcher.MustNewManaged("")
	l.NoSandbox(true)

	browser := rod.New().Client(l.Client()).MustConnect()
	
	defer browser.MustClose()

	page := browser.MustPage(url)
	
	page.MustWaitElementsMoreThan("div#jobs-list [style=\"\"]", 10)

	str, err := page.HTML()
	
	if err != nil {
		return "", err
	}
	
	return str, nil	
}

func MakeRequest(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, errors.New("Could not complete the request")
	}

	return res.Body, nil
}

func MakeRequestAsync(url string, ch chan<-[]byte) {
	res, err := http.Get(url)
	
	if (err != nil) {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	
	if (err != nil) {
		log.Fatal(err)
	}

	ch <- body
}
