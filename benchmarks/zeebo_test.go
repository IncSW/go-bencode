package benchmarks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeebo/bencode"
)

func Benchmark_Zeebo_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.EncodeBytes(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Zeebo_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent = nil
		err = bencode.DecodeBytes(unmarshalTestData, &torrent)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringInt64TestData, torrent)
}
