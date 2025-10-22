package taskManager

func NewUser(id, name string, email string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetName() string {
	return u.Name
}
