.PHONY: build
build:
	wire
	go build

.PHONY: run
run: build
	./sandbox

.PHONY: install-wire
install-wire:
	go install github.com/google/wire/cmd/wire@latest
