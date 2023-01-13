package main

import (
	"testing"
)

func Test_CalculateSumHash(t *testing.T) {
	tests := []struct {
		hash string
		want int
	}{
		{
			hash: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
			want: 841226,
		},
		{
			hash: "cb8379ac2098aa165029e3938a51da0bcecfc008fd6795f401178647f96c5b34",
			want: 401365080,
		},
		{
			hash: "4cd0e21a9a0795a14ec9aa5f0e7d1abff0492565770e43eafdf1e3e8afed1f33",
			want: 492566724,
		},
		{
			hash: "d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3",
			want: 10301339,
		},
	}

	for _, tt := range tests {
		t.Run(tt.hash, func(t *testing.T) {
			sum := <-calcSumHash(tt.hash)

			if tt.want != sum {
				t.Errorf("sum hash %s is %d, but got %d", tt.hash, tt.want, sum)
			}
		})
	}
}
