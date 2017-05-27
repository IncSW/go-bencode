package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/jackpal/bencode-go"
)

func BenchmarkJackpalBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Marshal(bytes.NewBuffer(nil), marshalTestData)
	}
}

func BenchmarkJackpalBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode.Decode(bytes.NewReader(unmarshalTestData))
	}
}
