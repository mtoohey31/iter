.PHONY: default ci test test-cov-check test-watch fmt fmt-check revive-check mdsh mdsh-check

default: test-watch

ci: fmt-check revive-check mdsh-check test-cov-check

fmt:
	gofmt -w .

fmt-check:
	test -z "$$(gofmt -l .)"

revive-check:
	revive -formatter friendly -set_exit_status -exclude ./testutils ./...

mdsh:
	mdsh

mdsh-check:
	mdsh --frozen

test:
	go test -coverprofile=/dev/null

test-cov-check:
	TEST_OUTPUT="$$(go test -coverprofile=/dev/null)" && echo "$$TEST_OUTPUT" && \
	  test "$$(echo "$$TEST_OUTPUT" | head -n2 | tail -n1 | awk '{ print $$2 }')" = '100.0%'

test-watch:
	gow -c test . -coverprofile=/dev/null
