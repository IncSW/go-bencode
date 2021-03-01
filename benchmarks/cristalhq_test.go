package benchmarks

import (
	"testing"

	"github.com/cristalhq/bencode"
)

func BenchmarkCristalhqMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Marshal(marshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCristalhqUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var torrent interface{}
		err := bencode.Unmarshal(unmarshalTestData, &torrent)
		if err != nil {
			b.Fatal(err)
		}
	}
}
