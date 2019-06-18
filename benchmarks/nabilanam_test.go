package benchmarks

import (
	"testing"

	"github.com/nabilanam/bencode/decoder"
	"github.com/nabilanam/bencode/encoder"
	"github.com/stretchr/testify/assert"
)

func BenchmarkNabilanamBencodeMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		encoder.New(marshalTestData).Encode()
	}
}

func BenchmarkNabilanamBencodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		decoder.New(unmarshalTestData).Decode()
	}
}

func TestNabilanamBencodeUnmarshal(t *testing.T) {
	assert := assert.New(t)
	nabilanam := decoder.New(unmarshalTestData).Decode()
	if !assert.Equal(nabilanam, marshalTestDataWithStrings) {
		return
	}
}
