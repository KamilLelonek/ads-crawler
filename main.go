package main

import "crawler/config"

func main() {
	router := config.SetupRouter()

	router.Run()
}
