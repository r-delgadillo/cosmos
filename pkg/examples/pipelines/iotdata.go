package pipelines

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/r-delgadillo/cosmos/pkg/examples/badger"
	"github.com/r-delgadillo/cosmos/pkg/examples/checkpoints"
)

// define a struct to represent IoT data
type IoTData struct {
	ID        string
	Timestamp time.Time
	Value     float64
}

func ProcessIotData(numMsg int) {
	// load the checkpoint
	var checkpoint checkpoints.Checkpoint
	checkpoints.Load(&checkpoint)

	// create a pipeline with three stages
	source := generator(numMsg) // generates 10 random IoT data points
	resumedData := make(chan IoTData)

	// resume processing from the last checkpoint
	go func() {
		fmt.Printf("Last checkpoint: %s\n", checkpoint.LastProcessedTime)
		for data := range source {
			if data.Timestamp.After(checkpoint.LastProcessedTime) {
				resumedData <- data
			} else {
				fmt.Printf("Message with ID '%s' has already been processed.\n", data.ID)
			}
		}
		close(resumedData)
	}()

	filtered := filter(resumedData) // filters out values < 0.5
	sink := average(filtered)       // calculates the average value of remaining data points

	// read the result from the pipeline and store in badger
	db := badger.NewClient()
	result := <-sink
	db.Update("avg", fmt.Sprintf("%f", result))
	db.View("avg")
	db.Close()

	// save the checkpoint
	checkpoint.LastProcessedTime = time.Now()
	checkpoints.Save(&checkpoint)

	// print the result
	fmt.Printf("Average value: %v\n", result)
}

// generate n random IoT data points and send them to a channel
func generator(n int) <-chan IoTData {
	out := make(chan IoTData)

	go func() {
		for i := 0; i < n; i++ {
			data := IoTData{
				ID:        strconv.FormatInt(int64(i), 10),
				Timestamp: time.Now().Add(-time.Second * 1),
				Value:     rand.Float64(),
			}
			out <- data
		}
		close(out)
	}()

	return out
}

// filter out data points with value < 0.5
func filter(in <-chan IoTData) <-chan IoTData {
	out := make(chan IoTData)

	go func() {
		for data := range in {
			if data.Value >= 0.5 {
				out <- data
			} else {
				fmt.Printf("Filtering device ID '%s'. Data value is less than 0.5. Actual '%v'\n", data.ID, data.Value)
			}
		}
		close(out)
	}()

	return out
}

// calculate the average value of the remaining data points
func average(in <-chan IoTData) <-chan float64 {
	out := make(chan float64)
	go func() {
		sum := 0.0
		count := 0

		for data := range in {
			sum += data.Value
			count++
		}

		if count > 0 {
			avg := sum / float64(count)
			out <- avg
		}

		close(out)
	}()
	return out
}
