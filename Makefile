coveralls:
	goveralls -service=travis-ci

test:
	go test -v -race

test-dict-encode:
	richgo test -v --count=1000 -run=TestMarshalUnOrderedDict

bench:
	go test -v -bench=. -run "^Benchmark"

all: test
