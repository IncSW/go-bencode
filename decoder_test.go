package bencode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	unmarshalTestData      = []byte("d8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment33:Debian CD from cdimage.debian.org4:infod6:lengthi170917888e4:name30:debian-8.8.0-arm64-netinst.iso12:piece lengthi262144eee")
	unmarshalWrongTestData = []byte("d4:infod6:lengthi170917888e12:piece lengthi262144e4:name30:debian-8.8.0-arm64-netinst.isoe8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment35:Debian CD from cdimage.debian.orge")
)

func TestUnmarshal(t *testing.T) {
	assert := assert.New(t)

	result, err := Unmarshal([]byte("i38e"))
	if !assert.NoError(err) || !assert.Equal(int64(38), result) {
		return
	}

	result, err = Unmarshal([]byte("38:udp://tracker.publicbt.com:80/announce"))
	if !assert.NoError(err) || !assert.Equal("udp://tracker.publicbt.com:80/announce", string(result.([]byte))) {
		return
	}

	result, err = Unmarshal([]byte("ll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee"))
	if !assert.NoError(err) || !assert.Equal([]interface{}{
		[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
		[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
	}, result) {
		return
	}

	result, err = Unmarshal([]byte("d8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceeee"))
	if !assert.NoError(err) || !assert.Equal(map[string]interface{}{
		"announce": []byte("udp://tracker.publicbt.com:80/announce"),
		"announce-list": []interface{}{
			[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
			[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
		},
	}, result) {
		return
	}

	_, err = Unmarshal(unmarshalTestData)
	if !assert.NoError(err) {
		return
	}

	_, err = Unmarshal(unmarshalWrongTestData)
	if !assert.Error(err) || !assert.Equal("bencode: not a valid bencoded string", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("i38"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: invalid integer field", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("i38qe"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: invalid integer byte: 113", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("l"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: invalid list field", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("d"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: invalid dictionary field", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("di38ee"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: non-string dictionary key", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("38"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: invalid string field", err.Error()) {
		return
	}

	result, err = Unmarshal([]byte("10:wasd"))
	if !assert.Error(err) || !assert.Nil(result) || !assert.Equal("bencode: not a valid bencoded string", err.Error()) {
		return
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Unmarshal(unmarshalTestData)
	}
}
