.PHONY: default ci test test-cov-check test-watch fmt fmt-check godoc-check

default: test-watch

ci: fmt-check godoc-check test-cov-check

fmt:
	gofmt -w .

fmt-check:
	test -z "$$(gofmt -l .)"

godoc-check:
	godoc-coverage .

test:
	go test -coverprofile=/dev/null

test-cov-check:
	TEST_OUTPUT="$$(go test -coverprofile=/dev/null)" && echo "$$TEST_OUTPUT" && \
	  test "$$(echo "$$TEST_OUTPUT" | head -n2 | tail -n1 | awk '{ print $$2 }' | grep -oP '^\d+')" -gt 95

test-watch:
	gow -c test . -coverprofile=/dev/null
