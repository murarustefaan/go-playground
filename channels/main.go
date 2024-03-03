package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type FetchDto struct {
	url      string
	message  []byte
	err      error
	duration time.Duration
}

func fetch(url string, channel chan<- FetchDto, limit chan bool) {
	limit <- true
	defer func() { <-limit }()

	start, duration := time.Now(), time.Duration(0)
	response, err := http.Get(url)

	if err != nil {
		duration = time.Since(start).Round(time.Millisecond)
		channel <- FetchDto{url, nil, err, duration}
		return
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		duration = time.Since(start).Round(time.Millisecond)
		channel <- FetchDto{url, nil, err, duration}
		return
	}

	duration = time.Since(start).Round(time.Millisecond)
	channel <- FetchDto{url, data, nil, duration}
}

func main() {
	start := time.Now()
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
	}

	channel := make(chan FetchDto)
	limit := make(chan bool, 2)
	
	for _, url := range urls {
		log.Printf("Fetching %-20s at %s\n", url, time.Now().Format(time.RFC3339Nano))
		go fetch(url, channel, limit)
	}

	for range urls {
		dto := <-channel
		log.Printf("Fetched %-20s in %s at %s\n", dto.url, dto.duration, time.Now().Format(time.RFC3339Nano))
		if dto.err != nil {
			log.Println(dto.err)
		}
	}
	log.Printf(" Done in %s\n", time.Since(start).Round(time.Millisecond))
}
