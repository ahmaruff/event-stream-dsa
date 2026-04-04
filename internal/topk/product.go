package topk

import (
	"fmt"

	"github.com/ahmaruff/event-stream-dsa/internal/model"
)

type Product struct {
	Event string
	Items map[int64]int64
}

type pair struct {
	key int64
	val int64
}

func (p *Product) Consume(e model.Event) error {
	if p.Items == nil {
		p.Items = make(map[int64]int64)
	}

	if e.Event == p.Event {
		p.Items[e.ItemId]++
	}

	return nil
}

func (p *Product) GetK() (int64, int64, error) {
	cnt := len(p.Items)

	if cnt < 1 {
		return 0, 0, fmt.Errorf("Product is empty")
	}

	items := make([]pair, 0, cnt)

	for k, v := range p.Items {
		items = append(items, pair{key: k, val: v})
	}

	if cnt > 1 {
		p.quickSort(items, 0, cnt-1)
	}

	top := items[0]
	return top.key, top.val, nil
}

func (p *Product) quickSort(arr []pair, low, high int) {
	if low < high {
		pivot := p.partition(arr, low, high)

		p.quickSort(arr, low, pivot-1)
		p.quickSort(arr, pivot+1, high)
	}
}

func (p *Product) partition(items []pair, low, high int) int {
	// 1. Cari index median
	pIdx := p.getPivotIndex(items, low, high)
	pivotVal := items[pIdx].val

	// 2. TUKER pivot ke ujung kanan (index high) biar aman gak keganggu loop
	p.swap(items, pIdx, high)

	// 3. i adalah batas "wilayah" angka yang lebih gede
	i := low - 1
	for j := low; j < high; j++ {
		if items[j].val > pivotVal {
			i++
			p.swap(items, i, j)
		}
	}

	// 4. Balikin pivot dari ujung (high) ke posisi tengah (i+1)
	p.swap(items, i+1, high)

	// return index posisi asli si pivot sekarang
	return i + 1
}

func (p *Product) swap(items []pair, i, j int) {
	items[i], items[j] = items[j], items[i]

}

func (p *Product) getPivotIndex(items []pair, low, high int) int {
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
