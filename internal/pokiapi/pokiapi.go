package main

import (
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func newClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
