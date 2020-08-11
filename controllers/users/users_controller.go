package users

import (
	"encoding/json"
	"fmt"
	"github.com/SHA056/bookstore_users/domain/users"
	"github.com/SHA056/bookstore_users/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	// Any request that comes through gets saved into the bytes object
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		//fmt.Println(err)
		return
	}
	fmt.Println("MARK 1>>>>>  reading bytes completed")

	// Using the bytes object to fill the details into another structure
	//json.Unmarshal(bytes, &user)

	if err := json.Unmarshal(bytes, &user) ; err != nil {
			// Handle error
			fmt.Println(err.Error())
			return
	}
	fmt.Println("MARK 2>>>>> Marshalling completed")
	//fmt.Println(string(bytes))
	result, saveError := services.CreateUser(user)
	fmt.Println("MARK 3>>>>> result saved")
	if saveError != nil {
		// Handle Error
		fmt.Println(saveError)
		return
	}
	fmt.Println("MARK 4>>>>> Returning JSON")
	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"not yet")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented,"not yet")
//}
//
