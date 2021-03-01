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
data, err := bencode.Marshal(value)
```

```go
data, err := bencode.Unmarshal(value)
```

## Performance

### Go 1.16, Debian 9.1, i7-7700

### Marshal

| Library             |    Time     | Bytes Allocated | Objects Allocated |
| :------------------ | :---------: | :-------------: | :---------------: |
| IncSW/go-bencode    | 803.3 ns/op |    176 B/op     |    6 allocs/op    |
| marksamman/bencode  | 828.4 ns/op |    384 B/op     |    8 allocs/op    |
| cristalhq/bencode   | 994.2 ns/op |    928 B/op     |    4 allocs/op    |
| aleksatr/go-bencode | 1048 ns/op  |    736 B/op     |    9 allocs/op    |
| nabilanam/bencode   | 2107 ns/op  |    1192 B/op    |   44 allocs/op    |
| jackpal/bencode-go  | 4766 ns/op  |    2016 B/op    |   45 allocs/op    |
| zeebo/bencode       | 4967 ns/op  |    1376 B/op    |   33 allocs/op    |

### Unmarshal

| Library             |    Time     | Bytes Allocated | Objects Allocated |
| :------------------ | :---------: | :-------------: | :---------------: |
| aleksatr/go-bencode | 768.5 ns/op |    640 B/op     |   16 allocs/op    |
| cristalhq/bencode   | 1160 ns/op  |    960 B/op     |   18 allocs/op    |
| nabilanam/bencode   | 1385 ns/op  |    1240 B/op    |   39 allocs/op    |
| IncSW/go-bencode    | 1386 ns/op  |    1128 B/op    |   25 allocs/op    |
| jackpal/bencode-go  | 2588 ns/op  |    1688 B/op    |   59 allocs/op    |
| marksamman/bencode  | 2827 ns/op  |    5768 B/op    |   54 allocs/op    |
| zeebo/bencode       | 6245 ns/op  |    6392 B/op    |   92 allocs/op    |

## License

[MIT License](LICENSE).
