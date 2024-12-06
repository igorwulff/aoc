package internal

import "time"

type BenchmarkResult struct {
	TotalTimeInMs int64
	Start         int64
}

func (b *BenchmarkResult) StartTimer() {
	b.Start = time.Now().UnixMilli()
}

func (b *BenchmarkResult) StopTimer() {
	b.TotalTimeInMs = time.Now().UnixMilli() - b.Start
}

func (b BenchmarkResult) GetTotalTimeInMs() int64 {
	return b.TotalTimeInMs
}
