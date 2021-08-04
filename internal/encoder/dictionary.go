package encoder

import (
	"sort"
	"sync"

	"github.com/IncSW/go-bencode/internal"
)

const stringsArrayLen = 20

var stringsArrayPool = sync.Pool{
	New: func() interface{} {
		return &[stringsArrayLen]string{}
	},
}

func sortStrings(ss []string) {
	if len(ss) <= stringsArrayLen {
		for i := 1; i < len(ss); i++ {
			for j := i; j > 0; j-- {
				if ss[j] >= ss[j-1] {
					break
				}
				ss[j], ss[j-1] = ss[j-1], ss[j]
			}
		}
	} else {
		sort.Strings(ss)
	}
}

//go:nosplit
func (e *Encoder) encodeDictionary(data map[string]interface{}) error {
	e.grow(1)
	e.writeByte('d')
	var keys []string
	if len(data) <= stringsArrayLen {
		stringsArray := stringsArrayPool.Get().(*[stringsArrayLen]string)
		defer stringsArrayPool.Put(stringsArray)
		keys = stringsArray[:0:len(data)]
	} else {
		keys = make([]string, 0, len(data))
	}
	for key, _ := range data {
		keys = append(keys, key)
	}
	sortStrings(keys)
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
