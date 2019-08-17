# ads-crawler

[![Build Status](https://travis-ci.org/KamilLelonek/ads-crawler.svg?branch=master)](https://travis-ci.org/KamilLelonek/ads-crawler)

Collecting `ads.txt` data from specified publishers.

## Authorized Digital Seller

Authorized Digital Seller list is a simple, flexible and secure method that publishers and distributors can use to publicly declare the companies they authorize to sell their digital inventory.

[`ads.txt`](https://iabtechlab.com/~iabtec5/wp-content/uploads/2016/07/IABOpenRTBAds.txtSpecification_Version1_Final.pdf) gives publishers control over their inventory in the market, making it harder for bad actors to profit from selling counterfeit inventory across the ecosystem.

### Use cases

1. Retrieve data from a specified publisher list (see below)
1. Validate data and ignore malformed or incomplete entries
1. Persist data in the PostgreSQL
1. Provide HTTP endpoint to retrieve dataset based on the publisher name

### List of Publishers

* http://www.cnn.com/ads.txt
* http://www.gizmodo.com/ads.txt
* http://www.nytimes.com/ads.txt
* https://www.bloomberg.com/ads.txt
* https://wordpress.com/ads.txt

## Installation

First of all, make sure you have [`Go 1.12` installed](https://golang.org/dl/).

Then, install all project dependencies by running:

    go get

## Usage

To start the server on you localhost, run:

    go run main.go

And visit http://localhost/v1/publisher/:name in your browser with a particular Publisher `name` from the list.

## Testing

To run the entire test suite, execute:

    go test ./...
