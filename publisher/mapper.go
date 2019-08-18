package publisher

import "fmt"

func url(domain string) string {
	return fmt.Sprint("https://", domain, ".com/ads.txt")
}

func FetchList(domain string) ([]Row, error) {
	url := url(domain)

	body, err := HttpGet(url)
	if err != nil {
		return nil, err
	}

	rows, err := NewList(body)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
