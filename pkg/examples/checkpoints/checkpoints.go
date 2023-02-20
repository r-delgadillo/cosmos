package checkpoints

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

const (
	path = "./bin/checkpoints/checkpoint.gob"
)

// define a struct to represent a checkpoint
type Checkpoint struct {
	LastProcessedTime time.Time
}

func ensureDirectory() {
	// check if the directory "example" exists
	if _, err := os.Stat("./bin/checkpoints"); os.IsNotExist(err) {
		// create a new directory named "example"
		err := os.Mkdir("./bin/checkpoints", 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

	}
}

// load the checkpoint from a file
func Load(checkpoint *Checkpoint) {
	ensureDirectory()
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(checkpoint)
		file.Close()
	}
	if err != nil {
		fmt.Println("Error loading checkpoint:", err)
	}
}

// save the checkpoint to a file
func Save(checkpoint *Checkpoint) {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(checkpoint)
		file.Close()
	}
	if err != nil {
		fmt.Println("Error saving checkpoint:", err)
	}
}
