package benchmarks

import (
	"testing"

	bencode "github.com/aleksatr/go-bencode"
)

func BenchmarkAleksatrMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Encode(marshalTestData)
	}
}

func BenchmarkAleksatrUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Decode(unmarshalTestData)
	}
}
