package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/pyro2927/plexbeat/beater"
)

func main() {
	err := beat.Run("plexbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
