package benchmarks

import (
	"bytes"
	"testing"

	"github.com/marksamman/bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Marksamman_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer = bencode.Encode(stringInt64TestData)
		if buffer == nil {
			b.Fatal("got nil")
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
	// WARN: ignoring unsupported types without errors
}

func Benchmark_Marksamman_Unmarshal(b *testing.B) {
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
