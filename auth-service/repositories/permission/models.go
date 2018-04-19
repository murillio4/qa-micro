package permission

import "gopkg.in/mgo.v2/bson"

// Permission permission db model
type Permission struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Name  string
	Roles []string
}

// Reset reset Permission
func (p *Permission) Reset() { *p = Permission{} }

// GetUUID UUID get function
func (p *Permission) GetID() string {
	if p != nil {
		return string(p.ID)
	}
	return ""
}

// GetName Name get function
func (p *Permission) GetName() string {
	if p != nil {
		return p.Name
	}
	return ""
}

// GetRoles Roles get funciton
func (p *Permission) GetRoles() []string {
	if p != nil {
		return p.Roles
	}
	return nil
}
