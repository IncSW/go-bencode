package benchmarks

import (
	"testing"

	"github.com/zeebo/bencode"
)

func BenchmarkZeeboBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.EncodeBytes(marshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkZeeboBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var torrent interface{}
		err := bencode.DecodeBytes(unmarshalTestData, &torrent)
		if err != nil {
			b.Fatal(err)
		}
	}
}
