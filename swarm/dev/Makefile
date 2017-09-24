.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/etherbanking/go-etherbanking/cmd/swarm
	go build -o bin/gebc     github.com/etherbanking/go-etherbanking/cmd/gebc
	go build -o bin/bootnode github.com/etherbanking/go-etherbanking/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/etherbanking/go-etherbanking/swarm/...
