package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

// ReadRequestData reads (and then closes) the entire request body.
// Returns non-nil error in case of failure.
func ReadRequestData(r *http.Request) ([]byte, error) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("failed to close request body: %v", err)
		}
	}()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
