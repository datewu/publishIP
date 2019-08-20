package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	for {
		resp, err := theLoop()
		if err != nil {
			log.Error().
				Str("url", *thirdPart).
				Err(err).
				Msg("http get failed")
		}
		fmt.Println(resp)
		time.Sleep(6 * time.Hour)
	}
}

func theLoop() (map[string]string, error) {
	ip, err := get()
	if err != nil {
		return nil, err
	}
	msg := map[string]string{
		"name":    *name,
		"ip":      ip,
		"version": GitCommit,
	}
	resp, err := post(msg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
