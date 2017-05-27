package bencode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var unmarshalTestData = []byte("d4:infod6:lengthi170917888e12:piece lengthi262144e4:name30:debian-8.8.0-arm64-netinst.isoe8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment33:Debian CD from cdimage.debian.orge")

func TestUnmarshal(t *testing.T) {
	assert := assert.New(t)

	data, length, ok := readUntil([]byte("38:udp://tracker.publicbt.com:80/announce"), ':')
	if !assert.True(ok) || !assert.Equal(2, length) || !assert.Equal([]byte("38"), data) {
		return
	}

	result, err := Unmarshal([]byte("i38e"))
	if !assert.NoError(err) || !assert.Equal(int64(38), result) {
		return
	}

	result, err = Unmarshal([]byte("38:udp://tracker.publicbt.com:80/announce"))
	if !assert.NoError(err) || !assert.Equal([]byte("udp://tracker.publicbt.com:80/announce"), result) {
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
}

func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Unmarshal(unmarshalTestData)
	}
}
