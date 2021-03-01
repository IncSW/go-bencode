# bencode

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Package implements Bencode encoding and decoding in Go.

## Features

* Simple API.
* Clean and tested code.
* Optimized for speed.
* Dependency-free.

## Install

Go version 1.15+

```
go get github.com/cristalhq/bencode
```

## Example

Marshaling into Bencode

```go
// data to process, most of the types are supported
var data interface{} = map[string]interface{}{
    "1":     42,
    "hello": "world",
    "foo":   []string{"bar", "baz"},
}

buf, err := bencode.Marshal(data)
checkErr(err)
fmt.Printf("marshaled: %s\n", string(buf))

// or via Encoder:
w := &bytes.Buffer{} // or any other io.Writer
err = bencode.NewEncoder(w).Encode(data)
checkErr(err)

// Output:
// marshaled: d1:1i42e3:fool3:bar3:baze5:hello5:worlde
```

Unmarshaling from Bencode

```go
var data interface{}

buf := []byte("li1ei42ee")

err := bencode.Unmarshal(buf, &data)
checkErr(err)

// or via Decoder:
r := bytes.NewBufferString("li1ei42ee") // or any other io.Reader
err = bencode.NewDecoder(r).Decode(&data)
checkErr(err)

fmt.Printf("unmarshaled: %v\n", data)

// Output:
// unmarshaled: [1 42]
```

See those examples: [example_test.go](https://github.com/cristalhq/bencode/blob/master/example_test.go).

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/bencode/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/bencode/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/bencode
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/bencode
[reportcard-img]: https://goreportcard.com/badge/cristalhq/bencode
[reportcard-url]: https://goreportcard.com/report/cristalhq/bencode
[coverage-img]: https://codecov.io/gh/cristalhq/bencode/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/bencode
