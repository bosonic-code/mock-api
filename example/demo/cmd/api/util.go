package main

import (
	"log"
	"net/url"
	"os"
)

func MustGetURL(key string) *url.URL {
	val := MustGet(key)
	res, err := url.Parse(val)
	if err != nil {
		log.Fatalf("Bad setting for environment variablbe  %v - value needs to be a valid URL: %v", key, err)
	}
	return res
}

func MustGet(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		log.Fatalf("Bad setting for environment variable %v - a value is required", key)
	}
	return val
}
