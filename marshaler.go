package bencode

import (
	"fmt"
	"strconv"
)

func Marshal(data interface{}) ([]byte, error) {
	result := make([]byte, 512) // WTF
	length, _, err := marshal(data, &result, 0, 512)
	if err != nil {
		return nil, err
	}

	return result[:length], nil
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
		offset, length = marshalBytes([]byte(value), result, offset, length)

		return offset, length, nil

	case []interface{}:
		return marshalList(value, result, offset, length)

	case map[string]interface{}:
		return marshalDictionary(value, result, offset, length)

	default:
		return 0, 0, fmt.Errorf("bencode: unsupported type: %T", value)
	}
}

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

	if rate > 1 {
		newResult := make([]byte, length*rate)
		copy(newResult, (*result)[:length])
		length *= rate
		*result = newResult
	}

	return length
}

func marshalInt(data int64, result *[]byte, offset int, length int) (int, int) {
	intBuffer := []byte(strconv.FormatInt(data, 10))
	intBufferLength := len(intBuffer)
	length = prepareBuffer(result, offset, length, intBufferLength+2)

	(*result)[offset] = 'i'
	offset++
	copy((*result)[offset:], intBuffer)
	offset += intBufferLength
	(*result)[offset] = 'e'
	offset++

	return offset, length
}

func marshalBytes(data []byte, result *[]byte, offset int, length int) (int, int) {
	dataLength := len(data)
	lengthBuffer := []byte(strconv.Itoa(dataLength))
	lengthBufferLength := len(lengthBuffer)
	length = prepareBuffer(result, offset, length, lengthBufferLength+1+dataLength)

	copy((*result)[offset:], lengthBuffer)
	offset += lengthBufferLength
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

	for key, data := range data {
		offset, length = marshalBytes([]byte(key), result, offset, length)
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
