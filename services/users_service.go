package services

import "github.com/SHA056/bookstore_users/domain/users"

//  Holds the entire business logic

func CreateUser(user users.User) (*users.User, error) {
	return &user,nil
}