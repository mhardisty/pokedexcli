package main

import "os"

func exitCallback(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
