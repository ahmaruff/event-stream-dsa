package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
)

func ParseFile(path string, processor func(model.Event) error) error {

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", path, err)
	}

	defer file.Close()

	return Stream(file, processor)
}

func Stream(r io.Reader, processor func(model.Event) error) error {
	reader := csv.NewReader(r)

	// skip header
	if _, err := reader.Read(); err != nil {
		return err
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		event, err := ParseRecord(record)
		if err != nil {
			return err
		}

		if err := processor(event); err != nil {
			return err
		}

	}

	return nil

}

func ParseRecord(record []string) (model.Event, error) {
	if len(record) < 4 {
		return model.Event{}, fmt.Errorf("Invalid Record")
	}

	timestamp, err := strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		return model.Event{}, err
	}

	visitorId, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		return model.Event{}, err
	}

	event := record[2]

	itemId, err := strconv.ParseInt(record[3], 10, 64)
	if err != nil {
		return model.Event{}, err
	}

	return model.Event{
		Timestamp: timestamp,
		VisitorId: visitorId,
		Event:     event,
		ItemId:    itemId,
	}, nil
}
