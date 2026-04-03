package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
	"github.com/ahmaruff/event-stream-dsa/internal/parser"
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

type Preview struct {
	Rows  []model.Event
	Limit int
}

func (p *Preview) Consume(e model.Event) error {
	if len(p.Rows) < p.Limit {
		p.Rows = append(p.Rows, e)
	}
	return nil
}

func (p *Preview) Print() {
	if len(p.Rows) == 0 {
		fmt.Println("No data to preview.")
		return
	}

	fmt.Printf("0a - Data Preview (First %d Rows)\n", len(p.Rows))

	// minwidth, tabwidth, padding, padchar, flags
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)

	// Header
	fmt.Fprintln(w, "No\tTimestamp\tVisitor ID\tEvent\tItem ID")
	fmt.Fprintln(w, "--\t---------\t----------\t-----\t-------")

	for i, row := range p.Rows {
		// Asumsi Timestamp lo itu detik (Unix). Kalau milidetik ganti UnixMilli
		tm := time.Unix(row.Timestamp, 0).Format("2006-01-02 15:04:05")

		fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%d\n",
			i+1, tm, row.VisitorId, row.Event, row.ItemId)
	}

	w.Flush()
	fmt.Println()
}

func main() {

	path := "dataset/dataset.csv"

	fmt.Println("Starting event stream processing...")
	start := time.Now()

	// consumers init
	preview := Preview{Limit: 3}
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

	fmt.Println("--- Data Preview (First 3 Rows) ---")
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
