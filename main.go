package main

import (
	"crawler/publisher"
)

func main() {
	router, group := SetupRouter()
	repo := &publisher.Repo{}

	publisher.MountRoutes(group)
	repo.Migrate()
	defer repo.CloseDB()

	router.Run()
}
