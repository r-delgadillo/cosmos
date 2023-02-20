package main

import (
	"math/rand"
	"time"

	"github.com/r-delgadillo/cosmos/lib/timer"
	"github.com/r-delgadillo/cosmos/pkg/examples/pipelines"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	t := timer.StopWatch{}
	// var test []int = []int{2, 3}
	// for i := 0; i < 2; i++ {
	// 	test = append(test, i)
	// }
	t.Start()
	pipelines.ProcessIotData(10)
	t.Stop()
	t.PrintResults()
}
