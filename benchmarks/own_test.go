package benchmarks

import (
	"testing"

	bencode "github.com/IncSW/go-bencode"
)

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Marshal(marshalTestData)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Unmarshal(unmarshalTestData)
	}
}
