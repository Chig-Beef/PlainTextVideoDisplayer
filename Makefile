all: build run

build:
		go build -o "dev.exe"

run:
		dev "test.ptv"
