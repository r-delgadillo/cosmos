package main

import (
	"fmt"

	"github.com/r-delgadillo/cosmos/pkg/logic"
	"github.com/r-delgadillo/cosmos/pkg/serving/app"
)

func main() {
	fmt.Println("Hello, world!")
	logic.DoSomething()
	server := app.NewServer()
	server.Run()
}
