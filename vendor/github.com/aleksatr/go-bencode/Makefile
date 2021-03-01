.DEFAULT: test

test:
	go test $$(go list ./...)

test_coverage:
	go test -coverprofile cp.log $$(go list ./...)
	go tool cover -html=cp.log
	rm cp.log

bench:
	go test -short -bench . -run ^a $$(go list ./...)
