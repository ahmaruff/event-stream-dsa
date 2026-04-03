package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
)

type EventConsumer interface {
	Consume(model.Event) error
}

func Stream(r io.Reader, consumers ...EventConsumer) error {
	reader := csv.NewReader(r)

	// skip header
	if _, err := reader.Read(); err != nil {
		return err
	}

	for i := 1; ; i++ {
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

		// DISPATCH
		for _, c := range consumers {
			if err := c.Consume(event); err != nil {
				return fmt.Errorf("consumer error at line %d: %w", i, err)
			}
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
