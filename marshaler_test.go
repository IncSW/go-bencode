package bencode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var marshalTestData = map[string]interface{}{
	"announce": []byte("udp://tracker.publicbt.com:80/announce"),
	"announce-list": []interface{}{
		[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
		[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
	},
	"comment": []byte("Debian CD from cdimage.debian.org"),
	"info": map[string]interface{}{
		"name":         []byte("debian-8.8.0-arm64-netinst.iso"),
		"length":       170917888,
		"piece length": 262144,
	},
}

func TestMarshal(t *testing.T) {
	assert := assert.New(t)

	result, err := Marshal(int64(38))
	if !assert.NoError(err) || !assert.Equal([]byte("i38e"), result) {
		return
	}

	result, err = Marshal([]byte("udp://tracker.publicbt.com:80/announce"))
	if !assert.NoError(err) || !assert.Equal([]byte("38:udp://tracker.publicbt.com:80/announce"), result) {
		return
	}

	result, err = Marshal([]interface{}{
		[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
		[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
	})
	if !assert.NoError(err) || !assert.Equal([]byte("ll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee"), result) {
		return
	}

	result, err = Marshal(map[string]interface{}{
		"announce": []byte("udp://tracker.publicbt.com:80/announce"),
		"announce-list": []interface{}{
			[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
			[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
		},
	})
	if !assert.NoError(err) || !assert.Equal([]byte("d8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceeee"), result) {
		return
	}

	result, err = Marshal(nil)
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: unsupported type: <nil>", err.Error()) {
		return
	}
}

func TestMarshalUnOrderedDict(t *testing.T) {
	assert := assert.New(t)
	unorderedEncodedData := []byte(`d7:ttg_tag32:8032a74ec22927a5bd6367537eafd87e7:privatei1e6:source20:[ttys3.bencode.test]e`)
	expectedReencodedData := []byte(`d7:privatei1e6:source20:[ttys3.bencode.test]7:ttg_tag32:8032a74ec22927a5bd6367537eafd87ee`)
	data, err := Unmarshal(unorderedEncodedData)
	assert.NoError(err)
	//t.Logf("decoded data: %#v", data)
	reencoded, err := Marshal(data)
	assert.NoError(err)
	assert.Equal(expectedReencodedData, reencoded)
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Marshal(marshalTestData)
	}
}
