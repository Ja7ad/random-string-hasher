package main

import "testing"

func Benchmark_RandomString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		randomString(i)
	}
}
