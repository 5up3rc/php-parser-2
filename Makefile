PHPFILE=example.php

all: compile run

build:
	find . -type f -iregex '.*\.go' -exec gofmt -l -s -w '{}' +
	go build

run: build
	./php-parser $(PHPFILE)

test:
	go test ./... --cover

bench:
	go test -benchmem -bench=. ./php5
	go test -benchmem -bench=. ./php7

compile: ./php5/php5.go ./php7/php7.go ./scanner/scanner.go
	rm -f y.output

./scanner/scanner.go: ./scanner/scanner.l
	golex -o $@ $<

./php5/php5.go: ./php5/php5.y
	goyacc -o $@ $<

./php7/php7.go: ./php7/php7.y
	goyacc -o $@ $<
