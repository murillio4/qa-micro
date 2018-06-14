package permission

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName               = "auth"
	permissionCollection = "permissions"
)

// Repository permission repository interface
type Repository interface {
	Create(*Permission) error
	GetByUUID(string) (*Permission, error)
	GetByName(string) (*Permission, error)
	GetByRole(string) ([]*Permission, error)
	GetByRoles([]string) ([]*Permission, error)
}

// PermissionsRepository Repo state
type PermissionsRepository struct {
	mgoSession *mgo.Session
}

// NewPermissionsRepository create a new permissions repository
func NewPermissionsRepository(mgoSession *mgo.Session) *PermissionsRepository {
	repo := &PermissionsRepository{mgoSession}

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}

	repo.collection().EnsureIndex(index)

	return repo
}

// Create create new permission
func (repo *PermissionsRepository) create(permission *Permission) error {
	c := repo.collection()
	return c.Insert(permission)
}

// GetByID Get permission by uuid
func (repo *PermissionsRepository) GetByID(id bson.ObjectId) (*Permission, error) {
	var permission *Permission

	c := repo.collection()
	err := c.FindId(id).One(&permission)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

// GetByName Get permission by name
func (repo *PermissionsRepository) GetByName(name string) (*Permission, error) {
	var permission *Permission

	c := repo.collection()
	err := c.Find(bson.M{
		"name": string(name),
	}).One(&permission)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

// GetByRole Get all permissions with role
func (repo *PermissionsRepository) GetByRole(role bson.ObjectId) ([]*Permission, error) {
	var permissions []*Permission

	c := repo.collection()
	err := c.Find(bson.M{
		"roles": bson.M{
			"$elemMatch": bson.M{
				"$eq": role,
			},
		},
	}).All(&permissions)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// GetByRoles Get all permissions with at least one of roles
func (repo *PermissionsRepository) GetByRoles(roles []string) ([]*Permission, error) {
	var permissions []*Permission

	c := repo.collection()
	err := c.Find(bson.M{
		"roles": bson.M{
			"$elemMatch": bson.M{
				"$in": roles,
			},
		},
	}).All(&permissions)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (repo *PermissionsRepository) collection() *mgo.Collection {
	return repo.mgoSession.DB(dbName).C(permissionCollection)
}
