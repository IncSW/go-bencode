# go-bencode [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

## Installation

`go get github.com/IncSW/go-bencode`

```go
import bencode "github.com/IncSW/go-bencode"
```

## Quick Start

```go
data, err := Marshal(value)
```

```go
data, err := Unmarshal(value)
```

## Performance

### Marshal

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| IncSW/go-bencode | 1493 ns/op | 554 B/op | 15 allocs/op |
| chihaya/bencode | 3614 ns/op | 1038 B/op | 64 allocs/op |
| jackpal/bencode-go | 8497 ns/op | 2289 B/op | 66 allocs/op |

### Unmarshal

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| IncSW/go-bencode | 3151 ns/op | 1360 B/op | 46 allocs/op |
| chihaya/bencode | 5281 ns/op | 5984 B/op | 66 allocs/op |
| jackpal/bencode-go | 6850 ns/op | 3073 B/op | 102 allocs/op |

## License

[MIT License](LICENSE).
