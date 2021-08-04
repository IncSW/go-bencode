package benchmarks

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	bencode "github.com/tumdum/bencoding"
)

func Benchmark_Tumdum_Marshal(b *testing.B) {
	b.ReportAllocs()
	var err error
	for n := 0; n < b.N; n++ {
		buffer, err = bencode.Marshal(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Tumdum_MarshalTo(b *testing.B) {
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

var tumdumTorrent map[string]interface{}

func Benchmark_Tumdum_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tumdumTorrent = map[string]interface{}{}
		err = bencode.Unmarshal(unmarshalTestData, &tumdumTorrent)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, stringInt64TestData, tumdumTorrent)
}
