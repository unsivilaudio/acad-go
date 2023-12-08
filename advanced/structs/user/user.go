package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	// ...
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// func (u *user) clearUserName() {
// 	u.firstName = ""
// 	u.lastName = ""
// }

func New(firstName, lastName, birthdate string) (*User, error) {
	// @todo ...DO VALIDATION
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("You must provide a first name, last name, and birthdate.")
	}

	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	// ...validate fields
	return Admin{
		email:    email,
		password: password,
		User:     User{firstName: "Max", lastName: "Schwarz", birthdate: "03/10/1989"},
	}
}
