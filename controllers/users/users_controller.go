package users

import (
	"fmt"
	"github.com/SHA056/bookstore_users/domain/users"
	"github.com/SHA056/bookstore_users/services"
	"github.com/SHA056/bookstore_users/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//fmt.Println(user)

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

	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status,getError)
		fmt.Println(getError)
		return
	}
	c.JSON(http.StatusOK,user)
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented,"not yet")Status
//}
//
