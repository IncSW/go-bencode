package benchmarks

import (
	"testing"

	bencode "github.com/IncSW/go-bencode"
)

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Marshal(marshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Unmarshal(unmarshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}
