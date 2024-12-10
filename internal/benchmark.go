package internal

import "time"

type BenchmarkResult struct {
	totalTime int64
	start     int64
}

func (b *BenchmarkResult) StartTimer() {
	b.start = time.Now().UnixMicro()
}

func (b *BenchmarkResult) StopTimer() {
	b.totalTime = time.Now().UnixMicro() - b.start
}

func (b BenchmarkResult) GetTotalTime() float64 {
	return float64(b.totalTime) / 1000
}
