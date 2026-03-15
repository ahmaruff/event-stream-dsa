package util

import (
	"fmt"
	"runtime"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %.2f MB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %.2f MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %.2f MB\n", bToMb(m.Sys))
}

func bToMb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}
