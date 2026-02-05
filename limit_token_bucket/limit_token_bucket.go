package main

import (
	"fmt"
	"time"
)

func main() {
	var fillInterval = time.Millisecond * 100
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)

	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for msg := range ticker.C {
			select {
			case tokenBucket <- struct{}{}:
				fmt.Println("product token:", msg)
			default:
				fmt.Println("no product token:", msg)
			}
			fmt.Println("current token cnt:", len(tokenBucket), msg)
		}

	}

	go fillToken()
	time.Sleep(time.Hour)
}
