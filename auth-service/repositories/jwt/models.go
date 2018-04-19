package jwt

import (
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	ID        string `json:",omitempty"`
	Email     string `json:",omitempty"`
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Picture   string `json:",omitempty"`

	Roles       map[string]string `json:",omitempty"`
	Permissions map[string]string `json:",omitempty"`
}

func (u *UserInfo) GetID() string {
	if u != nil {
		return u.ID
	}
	return ""
}

func (u *UserInfo) GetEmail() string {
	if u != nil {
		return u.Email
	}
	return ""
}

func (u *UserInfo) GetFirstName() string {
	if u != nil {
		return u.FirstName
	}
	return ""
}

func (u *UserInfo) GetLastName() string {
	if u != nil {
		return u.LastName
	}
	return ""
}

func (u *UserInfo) GetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func (u *UserInfo) GetPicture() string {
	if u != nil {
		return u.Picture
	}
	return ""
}

func (u *UserInfo) GetRoles() map[string]string {
	if u != nil {
		return u.Roles
	}
	return nil
}

func (u *UserInfo) GetPermissions() map[string]string {
	if u != nil {
		return u.Permissions
	}
	return nil
}

type BaseClaims struct {
	jwtGo.StandardClaims
	UserInfo UserInfo `json:",omitempty"`
}

// GetExpireAt get expire date of refresh claims
func (bc *BaseClaims) GetExpireAt() time.Time {
	return time.Unix(bc.ExpiresAt, 0)
}
