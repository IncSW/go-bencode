package benchmarks

import (
	"bytes"
	"testing"

	"github.com/lajide/bencode"
	"github.com/stretchr/testify/assert"
)

var lajideUnmarshalTestData = bencode.Dict{
	"announce": "udp://tracker.publicbt.com:80/announce",
	"announce-list": bencode.List{
		bencode.List{"udp://tracker.publicbt.com:80/announce"},
		bencode.List{"udp://tracker.openbittorrent.com:80/announce"},
	},
	"comment": "Debian CD from cdimage.debian.org",
	"info": bencode.Dict{
		"name":         "debian-8.8.0-arm64-netinst.iso",
		"length":       int64(170917888),
		"piece length": int64(262144),
	},
}

func Benchmark_Lajide_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.Marshal(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	// assert.Equal(b, string(unmarshalTestData), string(buffer))
	// BUG: Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics). http://bittorrent.org/beps/bep_0003.html#bencoding
}

func Benchmark_Lajide_MarshalTo(b *testing.B) {
	bytesBuffer = bytes.NewBuffer(make([]byte, 0, 512))
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bytesBuffer.Reset()
		err = bencode.NewEncoder(bytesBuffer).Encode(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	// assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
	// BUG: Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics). http://bittorrent.org/beps/bep_0003.html#bencoding
}

func Benchmark_Lajide_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	var err error
	for n := 0; n < b.N; n++ {
		torrent, err = bencode.NewDecoder(bytes.NewReader(unmarshalTestData)).Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, lajideUnmarshalTestData, torrent)
}
