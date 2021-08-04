package benchmarks

import (
	"testing"

	bencode "github.com/stints/bencode"
)

func Benchmark_Stints_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer = bencode.NewEncoder().Encode(stringInt64TestData)
		if buffer == nil {
			b.Fatal("is nil")
		}
	}
	b.StopTimer()
	// assert.Equal(b, string(unmarshalTestData), string(buffer))
	// BUG: []byte encoding as a list
	// BUG: Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics). http://bittorrent.org/beps/bep_0003.html#bencoding
}

func Benchmark_Stints_Unmarshal(b *testing.B) {
	b.Skip()
}
