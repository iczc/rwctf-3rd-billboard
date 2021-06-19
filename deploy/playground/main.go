package main

import (
	"github.com/iczc/billboard/playground/api"
	"github.com/iczc/billboard/playground/config"
)

func main() {
	api.NewServer(config.NewConfig()).Run()
}
