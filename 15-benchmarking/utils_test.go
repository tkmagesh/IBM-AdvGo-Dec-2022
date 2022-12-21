package utils

import "testing"

func Benchmark_isPrime_1_73(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime_1(73)
	}
}

func Benchmark_isPrime_2_73(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime_2(73)
	}
}
