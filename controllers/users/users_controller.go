package users

import (
	"fmt"
	"github.com/SHA056/bookstore_users/domain/users"
	"github.com/SHA056/bookstore_users/services"
	"github.com/SHA056/bookstore_users/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	//// Any request that comes through gets saved into the bytes object
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	// Handle error
	//	//fmt.Println(err)
	//	return
	//}
	//fmt.Println("MARK 1>>>>>  reading bytes completed")
	//
	//// Using the bytes object to fill the details into another structure
	////json.Unmarshal(bytes, &user)
	//
	//if err := json.Unmarshal(bytes, &user) ; err != nil {
	//		// Handle error
	//		fmt.Println(err.Error())
	//		return
	//}

	// Basically does all of the above in a single line of code
	if err := c.ShouldBindJSON(&user); err != nil {
		// Handle JSON error
		restErr := errors.NewBadRequestError("invalid  json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}

	fmt.Println("MARK 2>>>>> Marshalling completed")
	//fmt.Println(string(bytes))

	// Sending to service for processing
	fmt.Println("Sending to service for processing")
	result, saveError := services.CreateUser(user)
	fmt.Println("MARK 3>>>>> result saved")
	if saveError != nil {
		// Handle Error
		c.JSON(saveError.Status,saveError)
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
