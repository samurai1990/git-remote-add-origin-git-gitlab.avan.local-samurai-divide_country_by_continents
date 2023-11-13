.PHONY: build 


all:build

build:
	-rm sort-countries
	@go build -o sort-countries
