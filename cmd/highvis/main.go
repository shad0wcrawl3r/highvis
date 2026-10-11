package main

import (
	"log"

	"github.com/shad0wcrawl3r/highvis/internal/config"
)

var cfg config.Config

func init() {
	cfg = config.InitConfig()
}

func main() {
	log.Println(cfg)
}
