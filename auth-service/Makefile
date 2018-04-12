ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

run:
	docker-compose run auth-service

build: build-proto build-docker

build-proto:
	protoc -I. \
		--proto_path=$(ROOT_DIR) \
		--go_out=. \
		--micro_out=. \
		proto/auth/auth.proto

build-docker:
	docker-compose build auth-service