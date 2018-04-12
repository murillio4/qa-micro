ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

run:
	docker-compose run auth-service

build:
	protoc -I. \
		--proto_path=$(ROOT_DIR) \
		--go_out=. \
		--micro_out=. \
		proto/auth/auth.proto
	docker-compose build auth-service