package encoder

import (
	"sort"
	"strconv"
)

// Encoder struct
type Encoder struct {
	source interface{}
}

// New instantiates Encoder
func New(source interface{}) *Encoder {
	return &Encoder{source: source}
}

// Encode given source to bencode string
func (e *Encoder) Encode() string {
	switch e.source.(type) {
	case int:
		return e.encodeInt()
	case string:
		return e.encodeString()
	case []interface{}:
		return e.encodeList()
	case map[string]interface{}:
		return e.encodeDictionary()
	}
	return ""
}

func (e *Encoder) encodeInt() string {
	return "i" + strconv.Itoa(e.source.(int)) + "e"
}

func (e *Encoder) encodeString() string {
	s := e.source.(string)
	return strconv.Itoa(len(s)) + ":" + s
}

func (e *Encoder) encodeList() string {
	s := "l"
	l := e.source.([]interface{})
	for _, v := range l {
		e.source = v
		s += e.Encode()
	}
	return s + "e"
}

func (e *Encoder) encodeDictionary() string {
	s := "d"
	d := e.source.(map[string]interface{})
	keys := getSortedKeys(d)
	for _, k := range keys {
		e.source = k
		s += e.encodeString()
		e.source = d[k]
		s += e.Encode()
	}
	return s + "e"
}

func getSortedKeys(m map[string]interface{}) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
