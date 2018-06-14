
run:
	docker-compose run user-service

build: build-proto build-docker

build-proto:
	protoc \
		--proto_path=$(GOPATH)/src:. \
		--go_out=. \
		--micro_out=. \
		proto/user.proto

build-docker:
	docker-compose build user-service