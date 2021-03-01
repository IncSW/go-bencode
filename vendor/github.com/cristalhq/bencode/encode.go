package bencode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
)

// An Encoder writes Bencode values to an output stream.
type Encoder struct {
	w   io.Writer
	buf *bytes.Buffer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return NewEncoderWithBuffer(w, make([]byte, 0, 512))
}

// NewEncoderWithBuffer returns a new encoder that writes to w.
func NewEncoderWithBuffer(w io.Writer, buf []byte) *Encoder {
	return &Encoder{
		w:   w,
		buf: bytes.NewBuffer(buf),
	}
}

// Encode writes the Bencode encoding of v to the stream.
func (e *Encoder) Encode(v interface{}) error {
	e.buf.Reset()
	if err := e.marshal(v); err != nil {
		return fmt.Errorf("bencode: encode failed: %w", err)
	}
	_, err := e.w.Write(e.buf.Bytes())
	return err
}

func (e *Encoder) marshal(v interface{}) error {
	switch v := v.(type) {
	case []byte:
		e.marshalBytes(v)
	case string:
		e.marshalString(v)

	case map[string]interface{}:
		return e.marshalDictionary(v)

	case []interface{}:
		return e.marshalSlice(v)

	case int, int8, int16, int32, int64:
		e.marshalIntGen(v)
	case uint, uint8, uint16, uint32, uint64:
		e.marshalIntGen(v)

	case float32:
		e.marshalInt(int64(math.Float32bits(v)))
	case float64:
		e.marshalInt(int64(math.Float64bits(v)))

	case bool:
		var n int64
		if v {
			n = 1
		}
		e.marshalInt(n)

	case Marshaler:
		raw, err := v.MarshalBencode()
		if err != nil {
			return err
		}
		e.buf.Write(raw)

	default:
		return e.marshalReflect(reflect.ValueOf(v))
	}
	return nil
}

func (e *Encoder) writeInt(n int64) {
	var bs [20]byte // max_str_len( math.MaxInt64, math.MinInt64 ) base 10
	buf := strconv.AppendInt(bs[0:0], n, 10)
	e.buf.Write(buf)
}

func (e *Encoder) marshalBytes(b []byte) error {
	e.writeInt(int64(len(b)))
	e.buf.WriteByte(':')
	e.buf.Write(b)
	return nil
}

func (e *Encoder) marshalString(s string) error {
	e.writeInt(int64(len(s)))
	e.buf.WriteByte(':')
	e.buf.WriteString(s)
	return nil
}

func (e *Encoder) marshalIntGen(val interface{}) error {
	var num int64
	switch val := val.(type) {
	case int64:
		num = int64(val)
	case int32:
		num = int64(val)
	case int16:
		num = int64(val)
	case int8:
		num = int64(val)
	case int:
		num = int64(val)
	case uint64:
		num = int64(val)
	case uint32:
		num = int64(val)
	case uint16:
		num = int64(val)
	case uint8:
		num = int64(val)
	case uint:
		num = int64(val)
	default:
		return fmt.Errorf("unknown int type %T", val)
	}
	e.marshalInt(num)
	return nil
}

func (e *Encoder) marshalInt(num int64) error {
	e.buf.WriteByte('i')
	e.writeInt(num)
	e.buf.WriteByte('e')
	return nil
}

func (e *Encoder) marshalReflect(val reflect.Value) error {
	switch val.Kind() {
	case reflect.Slice:
		return e.marshalSliceReflect(val)
	case reflect.Array:
		return e.marshalArrayReflect(val)

	case reflect.Map:
		return e.marshalMap(val)

	case reflect.Struct:
		return e.marshalStruct(val)

	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return nil
		}
		return e.marshal(val.Elem().Interface())

	default:
		return fmt.Errorf("Unknown kind: %q", val)
	}
}

func (e *Encoder) marshalSliceReflect(val reflect.Value) error {
	elemKind := val.Type().Elem().Kind()
	if elemKind == reflect.Uint8 {
		return e.marshalBytes(val.Bytes())
	}
	return e.marshalList(val)
}

func (e *Encoder) marshalArrayReflect(val reflect.Value) error {
	elemKind := val.Type().Elem().Kind()
	if elemKind != reflect.Uint8 {
		return e.marshalList(val)
	}

	e.writeInt(int64(val.Len()))
	e.buf.WriteByte(':')

	for i := 0; i < val.Len(); i++ {
		v := byte(val.Index(i).Uint())
		e.buf.WriteByte(v)
	}
	return nil
}

func (e *Encoder) marshalList(val reflect.Value) error {
	if val.Len() == 0 {
		e.buf.WriteString("le")
		return nil
	}

	e.buf.WriteByte('l')
	for i := 0; i < val.Len(); i++ {
		if err := e.marshal(val.Index(i).Interface()); err != nil {
			return err
		}
	}
	e.buf.WriteByte('e')
	return nil
}

func (e *Encoder) marshalMap(val reflect.Value) error {
	rawKeys := val.MapKeys()
	if len(rawKeys) == 0 {
		e.buf.WriteString("de")
		return nil
	}

	keys := make([]string, len(rawKeys))

	for i, key := range rawKeys {
		if key.Kind() != reflect.String {
			return errors.New("Map can be marshaled only if keys are of type 'string'")
		}
		keys[i] = key.String()
	}

	sortStrings(keys)

	e.buf.WriteByte('d')
	for _, key := range keys {
		e.marshalString(key)

		value := val.MapIndex(reflect.ValueOf(key))
		if err := e.marshal(value.Interface()); err != nil {
			return err
		}
	}
	e.buf.WriteByte('e')
	return nil
}

func (e *Encoder) marshalStruct(val reflect.Value) error {
	return nil
}

func (e *Encoder) marshalDictionary(dict map[string]interface{}) error {
	if len(dict) == 0 {
		e.buf.WriteString("de")
		return nil
	}

	// less than `strSliceLen` keys in dict? - take from pool
	var keys []string
	if len(dict) <= strSliceLen {
		strArr := getStrArray()
		defer putStrArray(strArr)
		keys = strArr[:0:len(dict)]
	} else {
		keys = make([]string, 0, len(dict))
	}

	for key := range dict {
		keys = append(keys, key)
	}

	sortStrings(keys)

	e.buf.WriteByte('d')
	for _, key := range keys {
		e.marshalString(key)
		if err := e.marshal(dict[key]); err != nil {
			return err
		}
	}
	e.buf.WriteByte('e')
	return nil
}

func (e *Encoder) marshalSlice(v []interface{}) error {
	if len(v) == 0 {
		e.buf.WriteString("le")
		return nil
	}

	e.buf.WriteByte('l')
	for _, data := range v {
		if err := e.marshal(data); err != nil {
			return err
		}
	}
	e.buf.WriteByte('e')
	return nil
}
