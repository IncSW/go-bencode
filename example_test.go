package bencode_test

import (
	"fmt"
	"testing"

	"github.com/IncSW/go-bencode"
)

func TestExample(t *testing.T) {
	var dict interface{} = map[string]interface{}{
		"int":    123,
		"string": "Hello, World",
		"list":   []interface{}{"foo", "bar"},
	}
	data, err := bencode.Marshal(dict)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
