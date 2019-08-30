build:
	go build -o ./bin/mesh ./cmd/mesh/

clean:
	rm -rf ./bin

install:
	go install ./cmd/mesh/

tar: build
	tar -czvf ./bin/mesh-v0.0.4-osx.tar.gz ./bin/mesh