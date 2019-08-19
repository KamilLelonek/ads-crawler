package main

import (
	"crawler/config"
	"crawler/publisher"
)

func main() {
	router := config.SetupRouter(publisher.MountRoutes)
	db := config.SetupDB()

	publisher.Migrate(db)

	defer db.Close()
	router.Run()
}
