.PHONY: intm test bench

intm: 
	go build -o $@ ./cmd/$@

test:
	go test -v -race ./...

bench:
	cd ./internal/adapter/merger/ && go test -v -bench=. -benchmem -count=1
