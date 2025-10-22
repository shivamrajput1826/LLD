package creational

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
	Phone string
}

type UserBuilder struct {
	user User
}

func NewUserBuilder(name, email string) *UserBuilder {
	return &UserBuilder{user: User{Name: name, Email: email}}
}

func (b *UserBuilder) WithAge(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

func Builder() {
	user := NewUserBuilder("Chirag", "chiggywho@gmail.com").WithAge(27).WithPhone("7988401206").Build()
	fmt.Printf("%+v\n", user)
}
