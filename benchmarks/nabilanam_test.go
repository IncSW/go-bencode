package benchmarks

import (
	"testing"

	"github.com/nabilanam/bencode/decoder"
	"github.com/nabilanam/bencode/encoder"
	"github.com/stretchr/testify/assert"
)

var nabilanamResult string

func Benchmark_Nabilanam_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		nabilanamResult = encoder.New(stringIntTestData).Encode()
		if nabilanamResult == "" {
			b.Fatal("got empty")
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), nabilanamResult)
	// INFO: just 4 types supported
}

func Benchmark_Nabilanam_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent = decoder.New(unmarshalTestData).Decode()
		if torrent == nil {
			b.Fatal("got nil")
		}
	}
	b.StopTimer()
	assert.Equal(b, stringIntTestData, torrent)
}
