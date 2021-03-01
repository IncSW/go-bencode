package benchmarks

import (
	"testing"

	"github.com/cristalhq/bencode"
)

func BenchmarkCristalhqMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Marshal(marshalTestData)
	}
}

func BenchmarkCristalhqUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var torrent interface{}
		bencode.Unmarshal(unmarshalTestData, &torrent)
	}
}
