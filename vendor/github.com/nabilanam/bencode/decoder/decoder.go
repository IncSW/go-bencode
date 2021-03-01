package decoder

import (
	"bytes"
	"strconv"
)

// Decoder struct
type Decoder struct {
	source   []byte
	position int
	char     byte
}

// New instantiates Decoder
func New(source []byte) *Decoder {
	d := &Decoder{source: source}
	d.readChar()
	return d
}

// Decode given bencoded []byte
func (d *Decoder) Decode() interface{} {
	switch d.char {
	case 'd':
		return d.decodeDictionary()
	case 'l':
		return d.decodeList()
	case 'i':
		return d.decodeInt()
	default:
		if isDigit(d.char) {
			return d.decodeString()
		}
		return nil
	}
}

func (d *Decoder) readChar() {
	if d.position >= len(d.source) {
		d.char = 0
		return
	}
	d.char = d.source[d.position]
}

func (d *Decoder) incrementPosition(pos int) {
	d.position += pos
	d.readChar()
}

func (d *Decoder) decodeInt() int {
	e := d.position + bytes.IndexByte(d.source[d.position:], byte('e'))
	d.position++
	s := d.position
	d.position += e - s
	i, _ := strconv.Atoi(string(d.source[s:d.position]))
	d.incrementPosition(1)
	return i
}

func (d *Decoder) decodeString() string {
	colon := d.position + bytes.IndexByte(d.source[d.position:], byte(':'))
	length, _ := strconv.Atoi(string(d.source[d.position:colon]))
	d.incrementPosition(colon - d.position + length + 1)
	return string(d.source[colon+1 : d.position])
}

func (d *Decoder) decodeList() []interface{} {
	l := []interface{}{}
	d.incrementPosition(1)
	for d.char != 'e' {
		l = append(l, d.Decode())
	}
	d.incrementPosition(1)
	return l
}

func (d *Decoder) decodeDictionary() map[string]interface{} {
	dic := map[string]interface{}{}
	d.incrementPosition(1)
	for d.char != 'e' {
		k := d.decodeString()
		v := d.Decode()
		dic[k] = v
	}
	d.incrementPosition(1)
	return dic
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
