package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/cuberat/go-bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Cuberat_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bytesBuffer = bytes.NewBuffer(nil)
		err = bencode.Encode(bytesBuffer, stringInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
	// BUG: []byte encoding as a list
}

func Benchmark_Cuberat_MarshalTo(b *testing.B) {
	b.ReportAllocs()
	bytesBuffer = bytes.NewBuffer(make([]byte, 0, 512))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bytesBuffer.Reset()
		err = bencode.NewEncoder(bytesBuffer).Encode(stringInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
	// BUG: []byte encoding as a list
}

func Benchmark_Cuberat_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent, err = bencode.NewDecoder(bytes.NewReader(unmarshalTestData)).Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringInt64TestData, torrent)
}
