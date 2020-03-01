package requests

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// Get makes a get request to the given url and returns the body of the request.
func Get(url string) (string, error) {
	log.Println("[GET] " + url)
	resp, err := http.Get(url) // make get request
	if err != nil {            // check for failure
		if resp != nil {
			err2 := resp.Body.Close()
			if err2 != nil {
				return "", err2
			}
		}
		return "", err
	}

	// read from buffer to string, this is slow but safe
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	err2 := resp.Body.Close()

	// check for errors
	if err != nil {
		return "", err
	}

	if err2 != nil {
		return "", err2
	}

	return buf.String(), nil // return body of request
}

// Get makes a get request to the given url and returns the raw body of the request.
func GetRaw(url string) (io.Reader, error) {
	log.Println("[GET] " + url)
	resp, err := http.Get(url) // make get request
	if err != nil {            // check for failure
		return nil, err
	}
	body := resp.Body
	err = resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, nil // return body of request
}
