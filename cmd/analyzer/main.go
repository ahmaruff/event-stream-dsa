package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
	"github.com/ahmaruff/event-stream-dsa/internal/parser"
	"github.com/ahmaruff/event-stream-dsa/internal/util"
)

func main() {

	path := "dataset/dataset.csv"

	fmt.Println("Starting event stream processing...")
	util.PrintMemUsage()
	fmt.Println()

	count := 0
	start := time.Now()

	err := parser.ParseFile(path, func(e model.Event) error {

		count++

		if count%100000 == 0 {
			fmt.Printf("Processed: %d events\n", count)
			util.PrintMemUsage()
			fmt.Println()
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	eps := float64(count) / elapsed.Seconds()

	fmt.Println()
	fmt.Println("Finished processing dataset")
	fmt.Println()

	fmt.Printf("Events processed : %d\n", count)
	fmt.Printf("Elapsed time     : %.2f s\n", elapsed.Seconds())
	fmt.Printf("Throughput       : %.0f events/sec\n", eps)

	util.PrintMemUsage()
}
