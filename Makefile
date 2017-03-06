SRC=pce.go

all: build

build:
	go build

run:
	go run $(SRC)
