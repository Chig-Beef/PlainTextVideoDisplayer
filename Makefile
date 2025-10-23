all: build run

build:
		go build -o "dev.exe"
	
prod:
		go build -o "dev.exe" -ldflags "-w -s"

run:
		dev "test.ptv"
