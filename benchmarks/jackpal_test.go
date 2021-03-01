package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/jackpal/bencode-go"
)

func BenchmarkJackpalBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode.Marshal(bytes.NewBuffer(nil), marshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJackpalBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Decode(bytes.NewReader(unmarshalTestData))
		if err != nil {
			b.Fatal(err)
		}
	}
}
