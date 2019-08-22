build:
	go build -o ./bin/mesh ./cmd/mesh/
	go build -o ./bin/boot ./example/bootstrap/

clean:
	rm -rf ./bin

install:
	go install ./cmd/mesh/