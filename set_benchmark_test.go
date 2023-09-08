package main

import (
	"gosetgo/ds"
	v1 "gosetgo/v1"
	v2 "gosetgo/v2"
	v25 "gosetgo/v2.5"
	v3 "gosetgo/v3"
	v4 "gosetgo/v4"
	v5 "gosetgo/v5"
	"sync"
	"testing"
)

func BenchmarkV1Set(b *testing.B) {
	s := v1.NewSet()
	benchmarkSet(b, s)
}

// func BenchmarkV1SetParallel(b *testing.B) {
// 	s := v1.NewSet()
// 	benchmarkSetParallel(b, s)
// }

func BenchmarkV2Set(b *testing.B) {
	s := v2.NewSet()
	benchmarkSet(b, s)
}

func BenchmarkV2ßSetParallel(b *testing.B) {
	s := v2.NewSet()
	benchmarkSetParallel(b, s)
}

func BenchmarkV25Set(b *testing.B) {
	s := v25.NewSet()
	benchmarkSet(b, s)
}

func BenchmarkV25ßSetParallel(b *testing.B) {
	s := v25.NewSet()
	benchmarkSetParallel(b, s)
}

func BenchmarkV3Set(b *testing.B) {
	s := v3.NewSet()
	benchmarkSet(b, s)
}

func BenchmarkV3ßSetParallel(b *testing.B) {
	s := v3.NewSet()
	benchmarkSetParallel(b, s)
}

func BenchmarkV35Set(b *testing.B) {
	s := v3.NewSet()
	benchmarkSet(b, s)
}

func BenchmarkV35ßSetParallel(b *testing.B) {
	s := v3.NewSet()
	benchmarkSetParallel(b, s)
}

func BenchmarkV4Set(b *testing.B) {
	s := v4.NewSet()
	benchmarkSet(b, s)
}

func BenchmarkV4ßSetParallel(b *testing.B) {
	s := v4.NewSet()
	benchmarkSetParallel(b, s)
}

func BenchmarkV5Set(b *testing.B) {
	s := v5.NewSet()
	benchmarkSet(b, s)
}

func BenchmarkV5ßSetParallel(b *testing.B) {
	s := v5.NewSet()
	benchmarkSetParallel(b, s)
}

func benchmarkSet(b *testing.B, s ds.Set) {
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}

	for i := 0; i < b.N; i++ {
		s.Contains(i)
	}

	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}

	for i := 0; i < b.N; i++ {
		s.Contains(i)
	}
}

func benchmarkSetParallel(b *testing.B, s ds.Set) {
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Contains(i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Remove(i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Contains(i)
		}(i)
	}

	wg.Wait()
}
