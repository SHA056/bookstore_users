package app

import (
	"github.com/SHA056/bookstore_users/controllers/users"
	"github.com/SHA056/bookstore_users/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping) // Ping and not Ping() because we are defining, not executing the function
	router.GET("/users/:user_id", users.GetUser)


	router.POST("/users", users.CreateUser)
}
