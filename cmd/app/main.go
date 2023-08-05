package main

import (
	"github.com/r-delgadillo/cosmos/pkg/examples/consumer"
	"github.com/r-delgadillo/cosmos/pkg/examples/producer"
	"github.com/r-delgadillo/cosmos/pkg/serving/app"
)

func main() {
	// Create a channel for communication with the background goroutine
	ch := make(chan string)

	go producer.Send()
	go consumer.BackgroundRoutine(ch)
	server := app.NewServer()
	server.Run()
	for {

	}
}
