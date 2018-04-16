
all: build run

run:
	docker-compose run auth-service

build: build-proto build-docker

build-proto:
	protoc \
		--proto_path=$(GOPATH)/src:. \
		--go_out=. \
		--micro_out=. \
		proto/auth.proto

build-docker:
	docker-compose build auth-service