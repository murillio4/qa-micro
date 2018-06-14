package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Created time.Time
	Updated time.Time
	Deleted time.Time

	Email     string
	FirstName string
	LastName  string
	Name      string
	Picture   string

	Passwords []Password

	Roles []Role
}

func (u *User) GetID() bson.ObjectId {
	if u != nil {
		return u.ID
	}
	return ""
}

func (u *User) GetIDHex() string {
	if u != nil {
		return u.ID.Hex()
	}
	return ""
}

func (u *User) GetCreated() *time.Time {
	if u != nil {
		return &u.Created
	}
	return nil
}

func (u *User) GetUpdated() *time.Time {
	if u != nil {
		return &u.Updated
	}
	return nil
}

func (u *User) GetDeleted() *time.Time {
	if u != nil {
		return &u.Deleted
	}
	return nil
}

func (u *User) GetEmail() string {
	if u != nil {
		return u.Email
	}
	return ""
}

func (u *User) GetFirstName() string {
	if u != nil {
		return u.FirstName
	}
	return ""
}

func (u *User) GetLastName() string {
	if u != nil {
		return u.LastName
	}
	return ""
}

func (u *User) GetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func (u *User) GetPicture() string {
	if u != nil {
		return u.Picture
	}
	return ""
}

func (u *User) GetPasswords() []Password {
	if u != nil {
		return u.Passwords
	}
	return nil
}

func (u *User) GetRoles() []Role {
	if u != nil {
		return u.Roles
	}
	return nil
}

type Password struct {
	Created        time.Time
	PasswordString string
}

func (p *Password) GetPassword() string {
	if p != nil {
		return p.PasswordString
	}
	return ""
}

func (p *Password) GetCreated() *time.Time {
	if p != nil {
		return &p.Created
	}
	return nil
}

type Role struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Created time.Time
}

func (r *Role) GetID() bson.ObjectId {
	if r != nil {
		return r.ID
	}
	return ""
}

func (r *Role) GetIDHex() string {
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
