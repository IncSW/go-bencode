package benchmarks

import (
	"testing"

	bencode "github.com/owenliang/dht"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Owenliang_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.Encode(stringIntTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
	// INFO: just 4 types supported
}

func Benchmark_Owenliang_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent, err = bencode.Decode(unmarshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringIntTestData, torrent)
}
