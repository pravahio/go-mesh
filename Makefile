build:
	go build -o ./bin/mesh ./cmd/mesh/

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/mesh_amd64 ./cmd/mesh/

docker-build: build-linux
	docker build -t "pravahio/mesh:latest" .

clean:
	rm -rf ./bin

install:
	go install ./cmd/mesh/

run:
	./bin/mesh --rnz random --bs /ip4/192.168.0.103/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9 --rpc-cert-path ./docs/deploy/rpc.pravah.io.crt --rpc-key-path ./docs/deploy/rpc.pravah.io.key --auth-cert-path ./docs/deploy/auth.pravah.io.crt --debug --en-sub

tar: build build-linux
	tar -czvf ./bin/mesh-v0.1.0-osx.tar.gz ./bin/mesh
	rm  ./bin/mesh
	mv ./bin/mesh_amd64 ./bin/mesh
	tar -czvf ./bin/mesh-v0.1.0-amd64.tar.gz ./bin/mesh
	rm ./bin/mesh