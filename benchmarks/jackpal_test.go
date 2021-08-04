package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/jackpal/bencode-go"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Jackpal_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bytesBuffer = bytes.NewBuffer(nil)
		err = bencode.Marshal(bytesBuffer, bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
}

func Benchmark_Jackpal_MarshalTo(b *testing.B) {
	b.ReportAllocs()
	bytesBuffer = bytes.NewBuffer(make([]byte, 0, 512))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bytesBuffer.Reset()
		err = bencode.Marshal(bytesBuffer, bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
}

func Benchmark_Jackpal_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent, err = bencode.Decode(bytes.NewReader(unmarshalTestData))
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringInt64TestData, torrent)
}

func Benchmark_Jackpal_RealWorld(b *testing.B) {
	b.ReportAllocs()
	b.Run("Unmarshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bytesBuffer = bytes.NewBuffer(nil)
			torrent, err = bencode.Decode(bytes.NewReader(realWorldData))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("Marshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bytesBuffer = bytes.NewBuffer(nil)
			err = bencode.Marshal(bytesBuffer, torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	bytesBuffer = bytes.NewBuffer(buffer)
	b.Run("MarshalTo", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bytesBuffer.Reset()
			err = bencode.Marshal(bytesBuffer, torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	assert.Equal(b, string(realWorldData), bytesBuffer.String())
}
