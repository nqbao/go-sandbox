package sys

import (
	"log"
	"runtime"
)

// PrintMemUsage print current memory consumption of Go
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Printf(
		"Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v",
		(m.Alloc / 1024 / 1024),
		(m.TotalAlloc / 1024 / 1024),
		m.Sys,
		m.NumGC,
	)
}
