package users
// Defines objects that we are moving between the persistence layer and the application

import (
	"github.com/SHA056/bookstore_users/utils/errors"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	DateCreated string `json:"datecreated"`
}

func (user *User) Validate () *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}