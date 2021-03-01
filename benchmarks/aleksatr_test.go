package benchmarks

import (
	"testing"

	bencode "github.com/aleksatr/go-bencode"
)

func BenchmarkAleksatrMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Encode(marshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAleksatrUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode.Decode(unmarshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
}
