package benchmarks

import (
	"bytes"
	"testing"

	"github.com/chihaya/chihaya/frontend/http/bencode"
	"github.com/stretchr/testify/assert"
)

var chihayaUnmarshalTestData = bencode.Dict{
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

func Benchmark_Chihaya_Marshal(b *testing.B) {
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

func Benchmark_Chihaya_MarshalTo(b *testing.B) {
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

func Benchmark_Chihaya_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent, err = bencode.Unmarshal(unmarshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, chihayaUnmarshalTestData, torrent)
}

func Benchmark_Chihaya_RealWorld(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	b.Run("Unmarshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			torrent, err = bencode.Unmarshal(realWorldData)
			if err != nil {
				b.Fatal(err) // ERR: bencode: short read
			}
		}
	})
	b.Run("Marshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			buffer, err = bencode.Marshal(torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	bytesBuffer = bytes.NewBuffer(buffer)
	b.Run("MarshalTo", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bytesBuffer.Reset()
			err = bencode.NewEncoder(bytesBuffer).Encode(torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
