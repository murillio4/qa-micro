package role

import (
	"errors"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "user"
	roleCollection = "roles"
)

var (
	errIDNotHex = errors.New("Invalid id not type hex")
)

type Repository interface{}

type RoleRepository struct {
	mgoSession *mgo.Session
}

// NewRoleRepository create a new user repository
func NewRoleRepository(mgoSession *mgo.Session) *RoleRepository {
	repo := &RoleRepository{mgoSession}

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}

	if err := repo.collection().EnsureIndex(index); err != nil {
		log.WithError(err).Panic("Mongo index unique")
	}

	return repo
}

// CreateRole create a new role
func (repo *RoleRepository) CreateRole(role *Role) error {
	return repo.collection().Insert(role)
}

//============= GET FUNCTIONS =============//

// GetRoleByID get a role with id
func (repo *RoleRepository) GetRoleByID(id bson.ObjectId) (*Role, error) {
	var role *Role

	if err := repo.collection().FindId(id).One(&role); err != nil {
		return nil, err
	}

	return role, nil
}

// GetRolesByID get roles by ids
func (repo *RoleRepository) GetRolesByID(id []bson.ObjectId) ([]*Role, error) {
	var roles []*Role

	query := bson.M{
		"_id": bson.M{
			"$in": id,
		},
	}

	if err := repo.collection().Find(query).One(&roles); err != nil {
		return nil, err
	}

	return roles, nil
}

// GetRoleByName get a role with name
func (repo *RoleRepository) GetRoleByName(name string) (*Role, error) {
	var role *Role

	query := bson.M{
		"name": name,
	}

	if err := repo.collection().Find(query).One(&role); err != nil {
		return nil, err
	}

	return role, nil
}

func (repo *RoleRepository) collection() *mgo.Collection {
	return repo.mgoSession.DB(dbName).C(roleCollection)
}
