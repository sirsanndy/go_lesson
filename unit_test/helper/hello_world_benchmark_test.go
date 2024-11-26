package helper

import "testing"

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Sandy" + string(rune(i)))
	}
}
