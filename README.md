# go-bencode
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/travis/IncSW/go-bencode.svg?style=flat-square)](https://travis-ci.org/IncSW/go-bencode)
[![Coverage Status](https://img.shields.io/coveralls/IncSW/go-bencode/master.svg?style=flat-square)](https://coveralls.io/github/IncSW/go-bencode)
[![Go Report Card](https://goreportcard.com/badge/github.com/IncSW/go-bencode?style=flat-square)](https://goreportcard.com/report/github.com/IncSW/go-bencode)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/IncSW/go-bencode)

## Installation

`go get github.com/IncSW/go-bencode`

```go
import bencode "github.com/IncSW/go-bencode"
```

## Quick Start

```go
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

// Output:
// d3:inti123e4:listl3:foo3:bare6:string12:Hello, Worlde
```

```go
data, err := bencode.Unmarshal(value)
```

## Performance [benchmarks](https://github.com/IncSW/go-bencode/tree/benchmarks/benchmarks)

### Go 1.16, Debian 9.1, i7-7700

### Marshal

| Library                      |    Time     | Bytes Allocated | Objects Allocated |
| :--------------------------- | :---------: | :-------------: | :---------------: |
| IncSW/go-bencode [MarshalTo] | 590.8 ns/op |    112 B/op     |    2 allocs/op    |
| IncSW/go-bencode [Marshal]   | 676.1 ns/op |    624 B/op     |    3 allocs/op    |
| marksamman/bencode           | 820.3 ns/op |    384 B/op     |    8 allocs/op    |
| cristalhq/bencode            | 994.2 ns/op |    928 B/op     |    4 allocs/op    |
| aleksatr/go-bencode          | 1061 ns/op  |    736 B/op     |    9 allocs/op    |
| nabilanam/bencode            | 2103 ns/op  |    1192 B/op    |   44 allocs/op    |
| jackpal/bencode-go           | 4676 ns/op  |    2016 B/op    |   45 allocs/op    |
| zeebo/bencode                | 4889 ns/op  |    1376 B/op    |   33 allocs/op    |

### Unmarshal

| Library             |    Time    | Bytes Allocated | Objects Allocated |
| :------------------ | :--------: | :-------------: | :---------------: |
| IncSW/go-bencode    | 1149 ns/op |    960 B/op     |   18 allocs/op    |
| cristalhq/bencode   | 1160 ns/op |    960 B/op     |   18 allocs/op    |
| nabilanam/bencode   | 1379 ns/op |    1240 B/op    |   39 allocs/op    |
| aleksatr/go-bencode | 2270 ns/op |    1816 B/op    |   51 allocs/op    |
| jackpal/bencode-go  | 2577 ns/op |    1688 B/op    |   59 allocs/op    |
| marksamman/bencode  | 2725 ns/op |    5768 B/op    |   54 allocs/op    |
| zeebo/bencode       | 5988 ns/op |    6392 B/op    |   92 allocs/op    |

## License

[MIT License](LICENSE).
