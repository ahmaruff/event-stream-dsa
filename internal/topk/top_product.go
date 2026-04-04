package topk

import (
	"fmt"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
)

type TopProduct struct {
	Event    string
	Products map[int64]int64
}

type pair struct {
	key int64
	val int64
}

func (t *TopProduct) Consume(e model.Event) error {
	if t.Products == nil {
		t.Products = make(map[int64]int64)
	}

	if e.Event == t.Event {
		t.Products[e.ItemId]++
	}

	return nil
}

func (t *TopProduct) GetK() (int64, int64, error) {
	cnt := len(t.Products)

	if cnt < 1 {
		return 0, 0, fmt.Errorf("Product is empty")
	}

	items := make([]pair, 0, cnt)

	for k, v := range t.Products {
		items = append(items, pair{key: k, val: v})
	}

	if cnt > 1 {
		t.quickSort(items, 0, cnt-1)
	}

	top := items[0]
	return top.key, top.val, nil
}

func (t *TopProduct) quickSort(arr []pair, low, high int) {
	if low < high {
		p := t.partition(arr, low, high)

		t.quickSort(arr, low, p-1)
		t.quickSort(arr, p+1, high)
	}
}

func (t *TopProduct) partition(items []pair, low, high int) int {
	// 1. Cari index median
	pIdx := t.getPivotIndex(items, low, high)
	pivotVal := items[pIdx].val

	// 2. TUKER pivot ke ujung kanan (index high) biar aman gak keganggu loop
	t.swap(items, pIdx, high)

	// 3. i adalah batas "wilayah" angka yang lebih gede
	i := low - 1
	for j := low; j < high; j++ {
		if items[j].val > pivotVal {
			i++
			t.swap(items, i, j)
		}
	}

	// 4. Balikin pivot dari ujung (high) ke posisi tengah (i+1)
	t.swap(items, i+1, high)

	// return index posisi asli si pivot sekarang
	return i + 1
}

func (t *TopProduct) swap(items []pair, i, j int) {
	items[i], items[j] = items[j], items[i]

}

func (t *TopProduct) getPivotIndex(items []pair, low, high int) int {
	mid := low + (high-low)/2

	a := items[low].val
	b := items[mid].val
	c := items[high].val

	if (a < b && b < c) || (c < b && b < a) {
		return mid
	} else if (b < a && a < c) || (c < a && a < b) {
		return low
	} else {
		return high
	}
}
