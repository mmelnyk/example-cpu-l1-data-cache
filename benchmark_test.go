package sample

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkVersion0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Version0(1)
	}
}

func BenchmarkVersion1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Version1(1)
	}
}

func BenchmarkVersion1p(b *testing.B) {
	ncpu := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < ncpu; i++ {
		wg.Add(1)
		go func(N int) {
			for n := 0; n < N; n++ {
				Version1(1)
			}
			wg.Done()
		}(b.N / ncpu)
	}
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkVersion2(b *testing.B) {
	ncpu := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < ncpu; i++ {
		wg.Add(1)
		go func(N int) {
			for n := 0; n < N; n++ {
				Version2(1)
			}
			wg.Done()
		}(b.N / ncpu)
	}
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkVersion3(b *testing.B) {
	ncpu := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < ncpu; i++ {
		wg.Add(1)
		go func(N int) {
			for n := 0; n < N; n++ {
				Version3(1)
			}
			wg.Done()
		}(b.N / ncpu)
	}
	b.ResetTimer()
	wg.Wait()
}

func BenchmarkVersion4(b *testing.B) {
	ncpu := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < ncpu; i++ {
		wg.Add(1)
		go func(i, N int) {
			for n := 0; n < N; n++ {
				Version4(i, 1)
			}
			wg.Done()
		}(i, b.N/ncpu)
	}
	b.ResetTimer()
	wg.Wait()
}
