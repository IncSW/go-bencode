package benchmarks

import (
	"testing"

	bencode "github.com/chihaya/chihaya/frontend/http/bencode"
)

func BenchmarkChihayaBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Marshal(marshalTestData)
	}
}

func BenchmarkChihayaBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Unmarshal(unmarshalTestData)
	}
}
