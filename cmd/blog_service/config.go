package main

import "os"

type Config struct {
	ShortURLServiceHost string
}

func NewConfig() Config {
	shortURLServiceHost, ok := os.LookupEnv("SHORT_URL_SERVICE_HOST")
	if !ok {
		shortURLServiceHost = "https://api.s.url/"
	}
	return Config{ShortURLServiceHost: shortURLServiceHost}
}
