all : clean fmt build test install

clean :
	go clean

fmt :
	go fmt

build :
	go build

test :
	go test

install :
	go install