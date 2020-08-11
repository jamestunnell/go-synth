package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ReadRequestData(r *http.Request) ([]byte, error) {
	defer func(){
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
