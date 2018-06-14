package main

import (
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"

	pb "github.com/murillio4/qa-micro/auth-service/proto"
)

const name = "qa.micro.auth"

func main() {
	srv := micro.NewService(
		micro.Name(name),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterAuthServiceHandler(srv.Server(), new(pb.AuthService))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Failed to serve")
	}
}
