package benchmarks

import (
	"bytes"
	"testing"

	"github.com/marksamman/bencode"
)

func BenchmarkMarksammanBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Encode(marshalTestData)
	}
}

func BenchmarkMarksammanBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Decode(bytes.NewReader(unmarshalTestData))
		if err != nil {
			b.Fatal(err)
		}
	}
}
