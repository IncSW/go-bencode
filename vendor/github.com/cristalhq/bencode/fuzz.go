// +build gofuzz
//
// To run the fuzzer, run the following commands:
//		$ GO111MODULE=off go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
//		$ cd $GOPATH/src/github.com/cristalhq/bencode
//		$ go-fuzz-build
//		$ go-fuzz
// Note: go-fuzz doesn't support go modules, so you must have your local
// installation of bencode under $GOPATH.

package bencode

func Fuzz(data []byte) int {
	var dst interface{}

	if err := Unmarshal(data, &dst); err != nil {
		return 0
	}
	return 1
}
