package role

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Role struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`

	Created time.Time
	Updated time.Time
	Deleted time.Time

	Name string
}

func (r *Role) GetID() string {
	if r != nil {
		return r.ID.Hex()
	}
	return ""
}

func (r *Role) GetCreated() *time.Time {
	if r != nil {
		return &r.Created
	}
	return nil
}

func (r *Role) GetUpdated() *time.Time {
	if r != nil {
		return &r.Updated
	}
	return nil
}

func (r *Role) GetDeleted() *time.Time {
	if r != nil {
		return &r.Deleted
	}
	return nil
}

func (r *Role) GetName() string {
	if r != nil {
		return r.Name
	}
	return ""
}
