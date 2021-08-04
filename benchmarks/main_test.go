package benchmarks

import (
	"bytes"
	_ "embed"
)

var (
	//go:embed ubuntu-21.04-desktop-amd64.iso.torrent
	realWorldData []byte

	torrent     interface{}
	buffer      []byte
	bytesBuffer *bytes.Buffer
	err         error

	bytesInt64TestData = map[string]interface{}{
		"announce": []byte("udp://tracker.publicbt.com:80/announce"),
		"announce-list": []interface{}{
			[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
			[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
		},
		"comment": []byte("Debian CD from cdimage.debian.org"),
		"info": map[string]interface{}{
			"name":         []byte("debian-8.8.0-arm64-netinst.iso"),
			"length":       int64(170917888),
			"piece length": int64(262144),
		},
	}
	stringInt64TestData = map[string]interface{}{
		"announce": "udp://tracker.publicbt.com:80/announce",
		"announce-list": []interface{}{
			[]interface{}{"udp://tracker.publicbt.com:80/announce"},
			[]interface{}{"udp://tracker.openbittorrent.com:80/announce"},
		},
		"comment": "Debian CD from cdimage.debian.org",
		"info": map[string]interface{}{
			"name":         "debian-8.8.0-arm64-netinst.iso",
			"length":       int64(170917888),
			"piece length": int64(262144),
		},
	}
	stringIntTestData = map[string]interface{}{
		"announce": "udp://tracker.publicbt.com:80/announce",
		"announce-list": []interface{}{
			[]interface{}{"udp://tracker.publicbt.com:80/announce"},
			[]interface{}{"udp://tracker.openbittorrent.com:80/announce"},
		},
		"comment": "Debian CD from cdimage.debian.org",
		"info": map[string]interface{}{
			"name":         "debian-8.8.0-arm64-netinst.iso",
			"length":       170917888,
			"piece length": 262144,
		},
	}
	unmarshalTestData = []byte("d8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment33:Debian CD from cdimage.debian.org4:infod6:lengthi170917888e4:name30:debian-8.8.0-arm64-netinst.iso12:piece lengthi262144eee")
)
