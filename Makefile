all: test

test: bin/ocbcctl
	$< -alsologtostderr -h

bin/ocbcctl: $(shell find .)
	go build -o $@ .
	@echo "### BUILD COMPLETE"