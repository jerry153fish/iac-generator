package main

import (
	"github.com/jerry153fish/iac-generator/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLog()
	log.Info("Hello World")
}
