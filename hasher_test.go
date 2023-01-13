package main

import (
	"testing"
)

func Test_hashString(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{
			word: "hello",
			want: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			word: "world",
			want: "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7",
		},
		{
			word: "foo",
			want: "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae",
		},
		{
			word: "bar",
			want: "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9",
		},
		{
			word: "golang",
			want: "d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := newHash(tt.word); got.String() != tt.want {
				t.Errorf("newHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
