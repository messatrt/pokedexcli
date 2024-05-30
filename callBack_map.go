package main

import (
	"fmt"
)

func callBack_map(cfg *config) error {
	resp, err := cfg.pokiapiClient.ListLocationArea(cfg.nextUrl)
	if err != nil {
		return err
	}
	for _, area := range resp.Results {
		fmt.Printf("%s \n", area.Name)
	}
	cfg.nextUrl = resp.Next
	cfg.preUrl = resp.Previous
	return nil
}
