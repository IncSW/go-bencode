package benchmarks

import (
	"testing"

	bencode "github.com/aleksatr/go-bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Aleksatr_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.Encode(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Aleksatr_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent, err = bencode.Decode(unmarshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringInt64TestData, torrent)
}

func Benchmark_Aleksatr_RealWorld(b *testing.B) {
	b.ReportAllocs()
	b.Run("Unmarshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			torrent, err = bencode.Decode(realWorldData)
			if err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
	b.Run("Marshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			buffer, err = bencode.Encode(torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	assert.Equal(b, string(realWorldData), string(buffer))
}
