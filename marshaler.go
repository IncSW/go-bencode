package bencode

import "fmt"

func prepareBuffer(result *[]byte, offset int, length int, neededLength int) int {
	availableLength := length - offset
	if availableLength >= neededLength {
		return length
	}

	rate := 1
	for availableLength < neededLength {
		rate++
		availableLength = length*rate - offset
	}

	newResult := make([]byte, length*rate)
	copy(newResult, (*result)[:length])
	length *= rate
	*result = newResult

	return length
}

func Marshal(data interface{}) ([]byte, error) {
	return MarshalTo(make([]byte, 512), data)
}

func MarshalTo(dst []byte, data interface{}) ([]byte, error) {
	if cap(dst) > len(dst) {
		dst = dst[:cap(dst)]
	} else if len(dst) == 0 {
		dst = make([]byte, 512)
	}
	length, _, err := marshal(data, &dst, 0, len(dst))
	if err != nil {
		return nil, err
	}
	return dst[:length], nil
}

func marshal(data interface{}, result *[]byte, offset int, length int) (int, int, error) {
	switch value := data.(type) {
	case int64:
		offset, length = marshalInt(value, result, offset, length)
		return offset, length, nil

	case int32:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case int16:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case int8:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case int:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case uint64:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case uint32:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case uint16:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case uint8:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case uint:
		offset, length = marshalInt(int64(value), result, offset, length)
		return offset, length, nil

	case []byte:
		offset, length = marshalBytes(value, result, offset, length)
		return offset, length, nil

	case string:
		offset, length = marshalBytes(s2b(value), result, offset, length)
		return offset, length, nil

	case []interface{}:
		return marshalList(value, result, offset, length)

	case map[string]interface{}:
		return marshalDictionary(value, result, offset, length)

	default:
		return 0, 0, fmt.Errorf("bencode: unsupported type: %T", value)
	}
}

func marshalBytes(data []byte, result *[]byte, offset int, length int) (int, int) {
	dataLength := len(data)
	offset, length = writeInt(int64(dataLength), result, offset, length)
	length = prepareBuffer(result, offset, length, dataLength+1)
	(*result)[offset] = ':'
	offset++
	copy((*result)[offset:], data)
	offset += dataLength
	return offset, length
}

func marshalList(data []interface{}, result *[]byte, offset int, length int) (int, int, error) {
	length = prepareBuffer(result, offset, length, 1)

	(*result)[offset] = 'l'
	offset++

	for _, data := range data {
		var err error
		offset, length, err = marshal(data, result, offset, length)
		if err != nil {
			return 0, 0, err
		}
	}

	length = prepareBuffer(result, offset, length, 1)

	(*result)[offset] = 'e'
	offset++

	return offset, length, nil
}

func marshalDictionary(data map[string]interface{}, result *[]byte, offset int, length int) (int, int, error) {
	length = prepareBuffer(result, offset, length, 1)

	(*result)[offset] = 'd'
	offset++

	keys := make([]string, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}
	sortStrings(keys)

	for _, key := range keys {
		offset, length = marshalBytes(s2b(key), result, offset, length)
		var err error
		offset, length, err = marshal(data[key], result, offset, length)
		if err != nil {
			return 0, 0, err
		}
	}

	length = prepareBuffer(result, offset, length, 1)

	(*result)[offset] = 'e'
	offset++

	return offset, length, nil
}
