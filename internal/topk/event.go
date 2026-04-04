package topk

import (
	"fmt"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
)

type Events struct {
	Items map[string]int
}

func (events *Events) Consume(e model.Event) error {
	events.Items[e.Event]++
	return nil
}

func (events *Events) GetK() (string, int, error) {
	cnt := len(events.Items)
	if cnt < 1 {
		return "", 0, fmt.Errorf("Event is empty")
	}

	// 1. convert map to slice
	type pair struct {
		key string
		val int
	}

	items := make([]pair, 0, cnt)
	for k, v := range events.Items {
		items = append(items, pair{k, v})
	}

	// 2. Insertion Sort
	for i := 1; i < cnt; i++ {
		for j := i; j > 0; j-- {
			k := j - 1
			if items[k].val < items[j].val {
				// Swap manual
				items[k], items[j] = items[j], items[k]
			} else {
				break
			}
		}
	}

	top := items[0]
	return top.key, top.val, nil
}
