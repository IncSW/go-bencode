package benchmarks

import (
	"bytes"
	"testing"

	"github.com/lwch/bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Lwch_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bytesBuffer = bytes.NewBuffer(nil)
		err = bencode.NewEncoder(bytesBuffer).Encode(stringInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	// assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
	// BUG: []byte encoding as a list
	// BUG: Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics). http://bittorrent.org/beps/bep_0003.html#bencoding
}

func Benchmark_Lwch_MarshalTo(b *testing.B) {
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
	// assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
	// BUG: []byte encoding as a list
	// BUG: Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics). http://bittorrent.org/beps/bep_0003.html#bencoding
}

func Benchmark_Lwch_Unmarshal(b *testing.B) {
	b.Skip()
	torrent := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent = nil
		err = bencode.Decode(unmarshalTestData, &torrent)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, bytesInt64TestData, torrent)
	// BUG: not supported list in list
}
