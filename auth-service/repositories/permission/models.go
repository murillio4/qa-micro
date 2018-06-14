package permission

import "gopkg.in/mgo.v2/bson"

// Permission permission db model
type Permission struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string

	Roles []bson.ObjectId
}

// Reset reset Permission
func (p *Permission) Reset() { *p = Permission{} }

// GetUUID UUID get function
func (p *Permission) GetID() string {
	if p != nil {
		return p.ID.Hex()
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
func (p *Permission) GetRoles() []bson.ObjectId {
	if p != nil {
		return p.Roles
	}
	return nil
}

// GetRolesHex Roles get hex funciton
func (p *Permission) GetRolesHex() []string {
	if p != nil {
		hex := make([]string, len(p.Roles))
		for i := range p.Roles {
			hex[i] = p.Roles[i].Hex()
		}

		return hex
	}
	return nil
}
