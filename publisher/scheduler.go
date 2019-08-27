package publisher

import (
	"fmt"
	"log"
	"time"
)

const refreshPeriod = 1 * time.Minute

var ticker = time.NewTicker(refreshPeriod)
var publishers = []string{"cnn", "gizmodo", "nytimes", "bloomberg", "wordpress"}
var repo = &Repo{}

func Schedule() *time.Ticker {
	go func() {
		for range ticker.C {
			cachePublishers()
		}
	}()

	return ticker
}

func cachePublishers() {
	results := make(chan string)
	for _, publisher := range publishers {
		go cachePublisher(publisher, results)
	}

	for range publishers {
		fmt.Println(<-results)
	}
}

func cachePublisher(publisher string, result chan<- string) {
	rows, err := FetchList(publisher)

	if err != nil {
		log.Println(err)
	} else {
		storePublishers(publisher, rows)
	}
	result <- publisher
}

func storePublishers(publisher string, rows []Row) {
	for _, row := range rows {
		repo.WriteWithWebsite(publisher, row)
	}
}
