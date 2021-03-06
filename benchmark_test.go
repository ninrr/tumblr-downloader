package main

import (
	"strconv"
	"sync"
	"testing"
)

func benchmarkStrconvComparison(i int, b *testing.B) {
	var IDMutex sync.Mutex
	highestID := "122222222"

	if i == 0 {
		for n := 0; n < b.N; n++ {
			postIDint, _ := strconv.Atoi("110312919")

			IDMutex.Lock()
			highestIDint, _ := strconv.Atoi(highestID)
			if postIDint >= highestIDint {
				highestID = "110312919"
			}
			IDMutex.Unlock()
		}
	} else {
		b.SetParallelism(i)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				postIDint, _ := strconv.Atoi("110312919")

				IDMutex.Lock()
				highestIDint, _ := strconv.Atoi(highestID)
				if postIDint >= highestIDint {
					highestID = "110312919"
				}
				IDMutex.Unlock()
			}
		})
	}
}

func benchmarkInt64Comparison(i int, b *testing.B) {
	var IDMutex sync.Mutex
	highestID := int64(122222222)

	if i == 0 {
		for n := 0; n < b.N; n++ {
			idNew := int64(132145174)
			IDMutex.Lock()
			if highestID < idNew {
				highestID = idNew
			}
			IDMutex.Unlock()
		}
	} else {
		b.SetParallelism(i)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				idNew := int64(132145174)
				IDMutex.Lock()
				if highestID < idNew {
					highestID = idNew
				}
				IDMutex.Unlock()
			}
		})
	}

}

// === Benchmarks ===

func BenchmarkStrconvComparisonSingle(b *testing.B) {
	benchmarkStrconvComparison(0, b)
}

func BenchmarkStrconvComparisonParallel1(b *testing.B) {
	benchmarkStrconvComparison(1, b)
}

func BenchmarkStrconvComparisonParallel10(b *testing.B) {
	benchmarkStrconvComparison(10, b)
}

func BenchmarkInt64ComparisonSingle(b *testing.B) {
	benchmarkInt64Comparison(0, b)
}

func BenchmarkInt64ComparisonParallel1(b *testing.B) {
	benchmarkInt64Comparison(1, b)
}

func BenchmarkInt64ComparisonParallel10(b *testing.B) {
	benchmarkInt64Comparison(10, b)
}
