package main

import (
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"

	"github.com/murillio4/qa-micro/auth-service/connections"
	"github.com/murillio4/qa-micro/auth-service/repositories/permission"

	pb "github.com/murillio4/qa-micro/auth-service/proto"
)

const (
	name = "qa.micro.auth"
)

func main() {
	session, err := connections.CreateMongoSession()
	defer session.Close()

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panic("Could not connect to datastore with host")
	}

	repo, err := permission.NewPermissionsRepository(session)
	permissions, err := repo.GetRolePermissions([]string{"sdsd"})
	if err != nil {
		log.WithError(err).Error("Faen")
	}

	log.WithField("permissions", permissions[0]).Info("permissions")

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
