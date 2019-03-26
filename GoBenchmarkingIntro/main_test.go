package main

import (
	"testing"
)

func benchmarkCalculate(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
