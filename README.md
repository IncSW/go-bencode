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

## Performance [benchmarks](https://github.com/IncSW/go-bencode/tree/master/benchmarks)

### Go 1.16, Debian 9.1, i7-7700

### Marshal

| Library             |    Time     | Bytes Allocated | Objects Allocated | Notes |
| :------------------ | :---------: | :-------------: | :---------------: | :---: |
| IncSW/go-bencode    | 662.6 ns/op |    624 B/op     |    3 allocs/op    |       |
| cristalhq/bencode   | 939.4 ns/op |    928 B/op     |    4 allocs/op    |       |
| marksamman/bencode  | 1002 ns/op  |    736 B/op     |    9 allocs/op    |   4   |
| aleksatr/go-bencode | 1060 ns/op  |    736 B/op     |    9 allocs/op    |       |
| chihaya/chihaya     | 1722 ns/op  |    1009 B/op    |   53 allocs/op    |   1   |
| lajide/bencode      | 1725 ns/op  |    1011 B/op    |   53 allocs/op    |   1   |
| nabilanam/bencode   | 2865 ns/op  |    3192 B/op    |   54 allocs/op    |   5   |
| anacrolix/torrent   | 3179 ns/op  |    1328 B/op    |   25 allocs/op    |       |
| lwch/bencode        | 3340 ns/op  |    1792 B/op    |   75 allocs/op    | 1, 2  |
| tumdum/bencoding    | 3419 ns/op  |    1752 B/op    |   60 allocs/op    |       |
| stints/bencode      | 4018 ns/op  |    3120 B/op    |   100 allocs/op   | 1, 2  |
| ehmry/go-bencode    | 4569 ns/op  |    1496 B/op    |   33 allocs/op    |       |
| jackpal/bencode-go  | 4702 ns/op  |    2016 B/op    |   45 allocs/op    |       |
| zeebo/bencode       | 5003 ns/op  |    1376 B/op    |   33 allocs/op    |       |
| owenliang/dht       | 5180 ns/op  |    3279 B/op    |   80 allocs/op    |   5   |
| cuberat/go-bencode  | 5589 ns/op  |    1929 B/op    |   71 allocs/op    |   2   |

### MarshalTo

| Library            |    Time     | Bytes Allocated | Objects Allocated | Notes |
| :----------------- | :---------: | :-------------: | :---------------: | :---: |
| IncSW/go-bencode   | 581.0 ns/op |    112 B/op     |    2 allocs/op    |       |
| cristalhq/bencode  | 668.4 ns/op |     0 B/op      |    0 allocs/op    |       |
| chihaya/chihaya    | 1432 ns/op  |    307 B/op     |   49 allocs/op    |   1   |
| lajide/bencode     | 1462 ns/op  |    307 B/op     |   49 allocs/op    |   1   |
| anacrolix/torrent  | 2954 ns/op  |    720 B/op     |   21 allocs/op    |       |
| lwch/bencode       | 3093 ns/op  |    1089 B/op    |   71 allocs/op    | 1, 2  |
| tumdum/bencoding   | 3474 ns/op  |    1752 B/op    |   60 allocs/op    |       |
| jackpal/bencode-go | 4479 ns/op  |    1408 B/op    |   41 allocs/op    |       |
| ehmry/go-bencode   | 4650 ns/op  |    1528 B/op    |   33 allocs/op    |       |
| cuberat/go-bencode | 5360 ns/op  |    1321 B/op    |   67 allocs/op    |   2   |

### Unmarshal

| Library             |    Time     | Bytes Allocated | Objects Allocated | Notes |
| :------------------ | :---------: | :-------------: | :---------------: | :---: |
| IncSW/go-bencode    | 991.5 ns/op |    960 B/op     |   18 allocs/op    |       |
| cristalhq/bencode   | 1160 ns/op  |    960 B/op     |   18 allocs/op    |       |
| nabilanam/bencode   | 1379 ns/op  |    1240 B/op    |   39 allocs/op    |       |
| owenliang/dht       | 1702 ns/op  |    1352 B/op    |   46 allocs/op    |       |
| aleksatr/go-bencode | 2279 ns/op  |    1816 B/op    |   51 allocs/op    |       |
| jackpal/bencode-go  | 2597 ns/op  |    1688 B/op    |   59 allocs/op    |       |
| marksamman/bencode  | 2758 ns/op  |    5768 B/op    |   54 allocs/op    |       |
| ehmry/go-bencode    | 2865 ns/op  |    2064 B/op    |   41 allocs/op    |       |
| chihaya/chihaya     | 2961 ns/op  |    5880 B/op    |   61 allocs/op    |       |
| lajide/bencode      | 2973 ns/op  |    5880 B/op    |   61 allocs/op    |       |
| anacrolix/torrent   | 3723 ns/op  |    2456 B/op    |   62 allocs/op    |       |
| cuberat/go-bencode  | 4687 ns/op  |    6544 B/op    |   119 allocs/op   |       |
| zeebo/bencode       | 5954 ns/op  |    6376 B/op    |   91 allocs/op    |       |
| tumdum/bencoding    | 7891 ns/op  |    6568 B/op    |   157 allocs/op   |       |
| lwch/bencode        |      -      |        -        |         -         |   3   |
| stints/bencode      |      -      |        -        |         -         |   6   |

### RealWorld [ubuntu-21.04-desktop-amd64.iso.torrent](https://releases.ubuntu.com/21.04/ubuntu-21.04-desktop-amd64.iso.torrent)

| Library                       |    Time     | Bytes Allocated | Objects Allocated |
| :---------------------------- | :---------: | :-------------: | :---------------: |
| IncSW/go-bencode Unmarshal    | 1279 ns/op  |    1016 B/op    |   21 allocs/op    |
| IncSW/go-bencode Marshal      | 28572 ns/op |   262816 B/op   |    4 allocs/op    |
| IncSW/go-bencode MarshalTo    | 7800 ns/op  |    160 B/op     |    2 allocs/op    |
|                               |
|                               |
| cristalhq/bencode Unmarshal   | 1560 ns/op  |    1016 B/op    |   21 allocs/op    |
| cristalhq/bencode Marshal     | 49125 ns/op |   443168 B/op   |    5 allocs/op    |
| cristalhq/bencode MarshalTo   | 7709 ns/op  |     0 B/op      |    0 allocs/op    |
|                               |
|                               |
| aleksatr/go-bencode Unmarshal | 21615 ns/op |   223352 B/op   |   69 allocs/op    |
| aleksatr/go-bencode Marshal   | 26401 ns/op |   222689 B/op   |   13 allocs/op    |
|                               |
|                               |
| jackpal/bencode-go Unmarshal  | 48384 ns/op |   444484 B/op   |   81 allocs/op    |
| jackpal/bencode-go Marshal    | 55097 ns/op |   445694 B/op   |   60 allocs/op    |
| jackpal/bencode-go MarshalTo  | 35800 ns/op |   223128 B/op   |   54 allocs/op    |


#### Notes

1. BUG: Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics). http://bittorrent.org/beps/bep_0003.html#bencoding
2. BUG: []byte encoding as a list
3. BUG: not supported list in list
4. WARN: ignoring unsupported types without errors
5. INFO: just 4 types supported
6. INFO: files only

## License

[MIT License](LICENSE).
