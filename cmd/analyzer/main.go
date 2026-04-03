package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
	"github.com/ahmaruff/event-stream-dsa/internal/parser"
	"github.com/ahmaruff/event-stream-dsa/internal/preview"
	"github.com/ahmaruff/event-stream-dsa/internal/topk"
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
	start := time.Now()

	// consumers init
	preview := preview.Preview{Limit: 3}
	countEvents := CountEvents{Counts: 0}
	topK := topk.TopK{Events: map[string]int{}}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// pass consumer ke Stream
	err = parser.Stream(file, &preview, &countEvents, &topK)
	if err != nil {
		log.Fatal(err)
	}

	event, freq, err := topK.GetK()
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	eps := float64(countEvents.Counts) / elapsed.Seconds()

	fmt.Println()
	fmt.Println("Finished processing dataset")
	fmt.Println()

	preview.Print()
	fmt.Println("-----------------------------------")

	fmt.Println("0 - Stream Processing")
	fmt.Printf("Events processed : %d\n", countEvents.Counts)
	fmt.Printf("Elapsed time     : %.2f s\n", elapsed.Seconds())
	fmt.Printf("Throughput       : %.0f events/sec\n", eps)
	fmt.Println()
	fmt.Println()

	fmt.Println("0a - Top Event")
	fmt.Printf("Top Event        : %s\n", event)
	fmt.Printf("Frequency        : %d\n", freq)
	fmt.Println()
	fmt.Println()

	// fmt.Println("1 - Most Viewed Products")

	util.PrintMemUsage()
}
