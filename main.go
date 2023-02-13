package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jerry153fish/iac-generator/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error("Unexpected error", err)
		os.Exit(1)
	}
}
