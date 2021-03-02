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

func writeIntFirstBuffer(value uint32, result *[]byte, offset int, length int) (int, int) {
	length = prepareBuffer(result, offset, length, 3)
	start := value >> 24
	if start == 0 {
		(*result)[offset] = byte(value >> 16)
		offset++
		(*result)[offset] = byte(value >> 8)
		offset++
	} else if start == 1 {
		(*result)[offset] = byte(value >> 8)
		offset++
	}
	(*result)[offset] = byte(value)
	offset++
	return offset, length
}

func writeIntBuffer(value uint32, result *[]byte, offset int, length int) (int, int) {
	length = prepareBuffer(result, offset, length, 3)
	(*result)[offset] = byte(value >> 16)
	offset++
	(*result)[offset] = byte(value >> 8)
	offset++
	(*result)[offset] = byte(value)
	offset++
	return offset, length
}

func writeInt(value int64, result *[]byte, offset int, length int) (int, int) {
	if value < 0 {
		value = -value
		length = prepareBuffer(result, offset, length, 1)
		(*result)[offset] = '-'
		offset++
	}
	q1 := value / 1000
	if q1 == 0 {
		return writeIntFirstBuffer(digits[value], result, offset, length)
	}
	r1 := value - q1*1000
	q2 := q1 / 1000
	if q2 == 0 {
		offset, length = writeIntFirstBuffer(digits[q1], result, offset, length)
		return writeIntBuffer(digits[r1], result, offset, length)
	}
	r2 := q1 - q2*1000
	q3 := q2 / 1000
	if q3 == 0 {
		offset, length = writeIntFirstBuffer(digits[q2], result, offset, length)
		offset, length = writeIntBuffer(digits[r2], result, offset, length)
		return writeIntBuffer(digits[r1], result, offset, length)
	}
	r3 := q2 - q3*1000
	q4 := q3 / 1000
	if q4 == 0 {
		offset, length = writeIntFirstBuffer(digits[q3], result, offset, length)
		offset, length = writeIntBuffer(digits[r3], result, offset, length)
		offset, length = writeIntBuffer(digits[r2], result, offset, length)
		return writeIntBuffer(digits[r1], result, offset, length)
	}
	r4 := q3 - q4*1000
	q5 := q4 / 1000
	if q5 == 0 {
		offset, length = writeIntFirstBuffer(digits[q4], result, offset, length)
		offset, length = writeIntBuffer(digits[r4], result, offset, length)
		offset, length = writeIntBuffer(digits[r3], result, offset, length)
		offset, length = writeIntBuffer(digits[r2], result, offset, length)
		return writeIntBuffer(digits[r1], result, offset, length)
	}
	r5 := q4 - q5*1000
	q6 := q5 / 1000
	if q6 == 0 {
		offset, length = writeIntFirstBuffer(digits[q5], result, offset, length)
	} else {
		offset, length = writeIntFirstBuffer(digits[q6], result, offset, length)
		r6 := q5 - q6*1000
		offset, length = writeIntBuffer(digits[r6], result, offset, length)
	}
	offset, length = writeIntBuffer(digits[r5], result, offset, length)
	offset, length = writeIntBuffer(digits[r4], result, offset, length)
	offset, length = writeIntBuffer(digits[r3], result, offset, length)
	offset, length = writeIntBuffer(digits[r2], result, offset, length)
	return writeIntBuffer(digits[r1], result, offset, length)
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

func marshalInt(data int64, result *[]byte, offset int, length int) (int, int) {
	length = prepareBuffer(result, offset, length, 1)
	(*result)[offset] = 'i'
	offset++
	offset, length = writeInt(data, result, offset, length)
	length = prepareBuffer(result, offset, length, 1)
	(*result)[offset] = 'e'
	offset++
	return offset, length
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
