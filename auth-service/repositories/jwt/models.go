package jwt

type UserInfo struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Name      string
	Picture   string

	Roles       map[string]string
	Permissions map[string]string
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
