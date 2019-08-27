package main

import (
	"crawler/publisher"
)

func main() {
	router, group := SetupRouter()
	repo := &publisher.Repo{}
	scheduler := publisher.Schedule()

	publisher.MountRoutes(group)
	repo.Migrate()

	defer repo.CloseDB()
	defer scheduler.Stop()

	router.Run()
}
