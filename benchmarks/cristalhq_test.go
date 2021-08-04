package benchmarks

import (
	"testing"

	"github.com/cristalhq/bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Cristalhq_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.Marshal(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Cristalhq_MarshalTo(b *testing.B) {
	b.ReportAllocs()
	buffer = make([]byte, 0, 512)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.MarshalTo(buffer[:0], bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Cristalhq_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent = nil
		err = bencode.Unmarshal(unmarshalTestData, &torrent)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, bytesInt64TestData, torrent)
}

func Benchmark_Cristalhq_RealWorld(b *testing.B) {
	b.ReportAllocs()
	b.Run("Unmarshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			torrent = nil
			err = bencode.Unmarshal(realWorldData, &torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
	b.Run("Marshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			buffer, err = bencode.Marshal(torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("MarshalTo", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			buffer, err = bencode.MarshalTo(buffer[:0], torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	assert.Equal(b, string(realWorldData), string(buffer))
}
