package api

import (
	"log"
	"strconv"
)

var (
	allowedRates = []int{22050, 44100, 48000, 96000, 192000}
)

func getSrate(vars map[string]string) (int, bool) {
	srateStr, found := vars["srate"]
	if !found {
		log.Print("srate not given")
		return 0, false
	}

	srate, err := strconv.Atoi(srateStr)
	if err != nil {
		log.Printf("failed to parse srate value %s", srateStr)
		return 0, false
	}

	found = false

	for _, allowedRate := range allowedRates {
		if srate == allowedRate {
			found = true
			break
		}
	}

	if !found {
		log.Printf("srate value %d not one of allowed %v", srate, allowedRates)
		return 0, false
	}

	return srate, true
}
