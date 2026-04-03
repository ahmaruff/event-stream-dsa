package preview

import "fmt"
import "text/tabwriter"
import "github.com/ahmaruff/event-stream-dsa/internal/model"
import "time"
import "os"

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

	fmt.Printf("--- Data Preview (First %d Rows) ---\n", len(p.Rows))

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
