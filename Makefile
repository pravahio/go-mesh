build:
	go build -o ./bin/mesh ./cmd/mesh/

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/mesh_amd64 ./cmd/mesh/

docker-build: build-linux
	docker build -t "pravah-node:latest" .

clean:
	rm -rf ./bin

install:
	go install ./cmd/mesh/

tar: build build-linux
	tar -czvf ./bin/mesh-v0.0.5-osx.tar.gz ./bin/mesh
	rm  ./bin/mesh
	mv ./bin/mesh_amd64 ./bin/mesh
	tar -czvf ./bin/mesh-v0.0.5-amd64.tar.gz ./bin/mesh
	rm ./bin/mesh