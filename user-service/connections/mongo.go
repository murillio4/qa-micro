package connections

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

const (
	defaultHost = "localhost"
	defaultPort = "27017"
)

// CreateMongoSession creates a mongo db session. connection to host
func CreateMongoSession() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		host = defaultHost
	}

	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = defaultPort
	}

	session, err := mgo.Dial(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
