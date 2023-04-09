CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen_proto.sh 
run:
	go run cmd/main.go

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

swag-init:
	swag init -g api/api.go -o api/docs