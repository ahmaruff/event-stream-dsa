package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
	"github.com/ahmaruff/event-stream-dsa/internal/parser"
	"github.com/ahmaruff/event-stream-dsa/internal/util"
)

type CountEvents struct {
	Counts int
}

func (c *CountEvents) Consume(e model.Event) error {
	c.Counts++
	return nil
}

func main() {

	path := "dataset/dataset.csv"

	fmt.Println("Starting event stream processing...")

	countEvents := CountEvents{Counts: 0}

	start := time.Now()

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = parser.Stream(file, &countEvents)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	eps := float64(countEvents.Counts) / elapsed.Seconds()

	fmt.Println()
	fmt.Println("Finished processing dataset")
	fmt.Println()

	fmt.Printf("Events processed : %d\n", countEvents.Counts)
	fmt.Printf("Elapsed time     : %.2f s\n", elapsed.Seconds())
	fmt.Printf("Throughput       : %.0f events/sec\n", eps)

	util.PrintMemUsage()
}
