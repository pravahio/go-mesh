build:
	go build -o ./bin/mesh ./cmd/mesh/

clean:
	rm -rf ./bin

install:
	go install ./cmd/mesh/