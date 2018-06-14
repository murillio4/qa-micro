package user

import (
	"time"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "user"
	userCollection = "users"
)

// Repository repo interface
type Repository interface {
	Create(*User) *User

	GetByID(string) (*User, error)
	GetByEmail(string) (*User, error)
	GetByName(string) (*User, error)

	UpdateGeneral(string, *User) error
	UpdatePassword(string, *Password) error
	UpdateRoles(string, []Role) error

	DeleteRoles(string, []string) error
}

// UserRepository user repository
type UserRepository struct {
	mgoSession *mgo.Session
}

// NewUserRepository create a new user repository
func NewUserRepository(mgoSession *mgo.Session) *UserRepository {
	repo := &UserRepository{mgoSession}

	index := mgo.Index{
		Key:        []string{"name", "email"},
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

//============= CREATE FUNCTIONS =============//

// Create create a new user
func (repo *UserRepository) Create(user *User) error {
	user.ID = bson.NewObjectId()

	return repo.collection().Insert(user)
}

//============= GET FUNCTIONS =============//

// GetByID find user by id
func (repo *UserRepository) GetByID(id bson.ObjectId) (*User, error) {
	var user *User

	if err := repo.collection().FindId(id).One(&user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail find user email
func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user *User

	query := bson.M{"email": email}

	if err := repo.collection().Find(query).One(&user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetByName find user by name
func (repo *UserRepository) GetByName(name string) (*User, error) {
	var user *User

	if err := repo.collection().Find(bson.M{"name": name}).One(&user); err != nil {
		return nil, err
	}

	return user, nil
}

//============= UPDATE FUNCTIONS =============//

// UpdateGeneral update email firstname lastname and picture
func (repo *UserRepository) UpdateGeneral(id bson.ObjectId, user *User) error {
	change := mgo.Change{
		Update: bson.M{
			"$currentDate": bson.M{
				"updated": true,
			},
			"$set": bson.M{
				"email":     user.GetEmail(),
				"firstname": user.GetFirstName(),
				"lastname":  user.GetLastName(),
				"picture":   user.GetPicture(),
			},
		},
		ReturnNew: false,
		Upsert:    false,
		Remove:    false,
	}

	_, err := repo.collection().FindId(id).Apply(change, &user)

	return err
}

// UpdatePassword updates password on user with id
func (repo *UserRepository) UpdatePassword(id bson.ObjectId, password Password) error {
	change := bson.M{
		"$push": bson.M{
			"passwords": password,
		},
	}

	return repo.collection().UpdateId(id, change)
}

// UpdateRoles appends a number of roles to an
func (repo *UserRepository) UpdateRoles(id bson.ObjectId, roles []Role) error {
	now := time.Now()

	for i := range roles {
		roles[i].Created = now
	}

	change := bson.M{
		"$addToSet": bson.M{
			"roles": bson.M{
				"$each": roles,
			},
		},
	}

	return repo.collection().UpdateId(id, change)
}

//============= DELETE FUNCTIONS =============//

// DeleteRoles removes roles from a user
func (repo *UserRepository) DeleteRoles(id string, roles []bson.ObjectId) error {
	change := bson.M{
		"$pullAll": bson.M{
			"roles": bson.M{
				"_id": bson.M{
					"$in": roles,
				},
			},
		},
	}

	return repo.collection().UpdateId(id, change)
}

func (repo *UserRepository) collection() *mgo.Collection {
	return repo.mgoSession.DB(dbName).C(userCollection)
}
