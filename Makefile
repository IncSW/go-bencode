coveralls:
	goveralls -service=travis-ci

test:
	go test -v -race

bench:
	go test -v -bench=. -run "^Benchmark"

all: test
