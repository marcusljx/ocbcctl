all: bin/ocbcctl

GCP_SA = $(wildcard ./.firebase/expenses*.json)

test: bin/ocbcctl
	./bin/ocbcctl login

bin/ocbcctl: tidy $(shell find *.go)
	go build -o $@ .
	@echo "### BUILD COMPLETE"

tidy:
	go mod tidy