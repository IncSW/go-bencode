package encoder

import "github.com/IncSW/go-bencode/internal"

func (e *Encoder) encodeDictionary(data map[string]interface{}) error {
	e.grow(1)
	e.writeByte('d')
	keys := make([]string, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}
	internal.SortStrings(keys)
	for _, key := range keys {
		e.encodeBytes(internal.S2B(key))
		err := e.encode(data[key])
		if err != nil {
			return err
		}
	}
	e.grow(1)
	e.writeByte('e')
	return nil
}
