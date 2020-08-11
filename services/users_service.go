package services

import (
	"github.com/SHA056/bookstore_users/domain/users"
	"github.com/SHA056/bookstore_users/utils/errors"
)

//  Holds the entire business logic

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user,nil
}