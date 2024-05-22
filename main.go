package main

import (
	"Gin/src/config"
	"Gin/src/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	router.Init()
}
