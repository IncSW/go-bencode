package bencode

import (
	"bytes"
	"errors"
	"sort"
	"strconv"
)

var ErrUnsupportedType = errors.New("unsupported type")

// Encode accepts either string/[]byte, int/int8/int16/int32/int64, []interface{} or map[string]interface{},
// interface{} being anything from mentioned types
// TODO: consider using io.Writer and writing directly to it
func Encode(value interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := encodeObject(buf, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func encodeObject(buf *bytes.Buffer, i interface{}) error {
	switch val := i.(type) {
	case int:
		return encodeInt(buf, int64(val))
	case int8:
		return encodeInt(buf, int64(val))
	case int16:
		return encodeInt(buf, int64(val))
	case int32:
		return encodeInt(buf, int64(val))
	case int64:
		return encodeInt(buf, val)
	case []byte:
		return encodeBytes(buf, val)
	case string:
		return encodeString(buf, val)
	case []interface{}:
		return encodeList(buf, val)
	case map[string]interface{}:
		return encodeDictionary(buf, val)
	}

	return ErrUnsupportedType
}

func encodeInt(buf *bytes.Buffer, i int64) error {
	// i
	err := buf.WriteByte(IntegerStart)
	if err != nil {
		return err
	}

	// actual digits
	_, err = buf.WriteString(strconv.FormatInt(i, 10))
	if err != nil {
		return err
	}

	// e
	return buf.WriteByte(IntegerEnd)
}

func encodeString(buf *bytes.Buffer, s string) error {
	// length
	_, err := buf.WriteString(strconv.Itoa(len(s)))
	if err != nil {
		return err
	}

	// :
	_, err = buf.Write(StringDelimiter)
	if err != nil {
		return err
	}

	// actual string
	_, err = buf.WriteString(s)
	return err
}

// TODO: try to merge with encodeString
// but without conversion []byte(str)
func encodeBytes(buf *bytes.Buffer, b []byte) error {
	// length
	_, err := buf.WriteString(strconv.Itoa(len(b)))
	if err != nil {
		return err
	}

	// :
	_, err = buf.Write(StringDelimiter)
	if err != nil {
		return err
	}

	// actual bytes
	_, err = buf.Write(b)
	return err
}

func encodeList(buf *bytes.Buffer, l []interface{}) error {
	// l
	err := buf.WriteByte(ListStart)
	if err != nil {
		return err
	}

	// actual members
	for _, m := range l {
		err = encodeObject(buf, m)
		if err != nil {
			return err
		}
	}

	// e
	return buf.WriteByte(ListEnd)
}

func encodeDictionary(buf *bytes.Buffer, d map[string]interface{}) error {
	// d
	err := buf.WriteByte(DictionaryStart)
	if err != nil {
		return err
	}

	// gather all keys
	keys := make([]string, len(d))
	i := 0
	for k := range d {
		keys[i] = k
		i++
	}

	// order them
	sort.Strings(keys)

	// encode KVs
	for _, k := range keys {
		err = encodeString(buf, k)
		if err != nil {
			return err
		}

		err = encodeObject(buf, d[k])
		if err != nil {
			return err
		}
	}

	// e
	return buf.WriteByte(DictionaryEnd)
}
