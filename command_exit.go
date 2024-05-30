package main

import "os"

func callBack_exit(cfg *config) error {
	os.Exit(0)
	return nil
}
