package benchmarks

import (
	"bytes"
	"testing"

	"github.com/anacrolix/torrent/bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Anacrolix_Marshal(b *testing.B) {
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

func Benchmark_Anacrolix_MarshalTo(b *testing.B) {
	b.ReportAllocs()
	bytesBuffer = bytes.NewBuffer(make([]byte, 0, 512))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bytesBuffer.Reset()
		err = bencode.NewEncoder(bytesBuffer).Encode(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
}

func Benchmark_Anacrolix_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent = nil
		err = bencode.Unmarshal(unmarshalTestData, &torrent)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringInt64TestData, torrent)
}
