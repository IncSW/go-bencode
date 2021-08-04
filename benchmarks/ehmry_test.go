package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/ehmry/go-bencode"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Ehmry_Marshal(b *testing.B) {
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

func Benchmark_Ehmry_MarshalTo(b *testing.B) {
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
	assert.Equal(b, string(unmarshalTestData), bytesBuffer.String())
}

var ehmryTorrent map[string]interface{}

func Benchmark_Ehmry_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		ehmryTorrent = nil
		err = bencode.Unmarshal(unmarshalTestData, &ehmryTorrent)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, bytesInt64TestData, ehmryTorrent)
}
