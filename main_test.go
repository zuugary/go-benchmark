package main_test

import (
	"testing"

	ex "github.com/zuugary/go-benchmark"
)

func benchmarkFact(b *testing.B, i int) {
	// いろいろ重い処理
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ex.Fact(i)

		b.StopTimer()
		// 別の処理
		b.StartTimer()
	}
}

func BenchmarkFact_10(b *testing.B)   { benchmarkFact(b, 10) }
func BenchmarkFact_20(b *testing.B)   { benchmarkFact(b, 20) }
func BenchmarkFact_50(b *testing.B)   { benchmarkFact(b, 50) }
func BenchmarkFact_100(b *testing.B)  { benchmarkFact(b, 100) }
func BenchmarkFact_1000(b *testing.B) { benchmarkFact(b, 1000) }

func benchmarkFact2(b *testing.B, i int) {
	for n := 0; n < b.N; n++ {
		ex.Fact2(i)
	}
}

func BenchmarkFact2_10(b *testing.B)   { benchmarkFact2(b, 10) }
func BenchmarkFact2_20(b *testing.B)   { benchmarkFact2(b, 20) }
func BenchmarkFact2_50(b *testing.B)   { benchmarkFact2(b, 50) }
func BenchmarkFact2_100(b *testing.B)  { benchmarkFact2(b, 100) }
func BenchmarkFact2_1000(b *testing.B) { benchmarkFact2(b, 1000) }
