package benchmarks

import (
	"testing"

	bencode "github.com/zeebo/bencode"
)

func BenchmarkZeeboBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.EncodeBytes(marshalTestData)
	}
}

func BenchmarkZeeboBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var torrent interface{}
		bencode.DecodeBytes(unmarshalTestData, &torrent)
	}
}
