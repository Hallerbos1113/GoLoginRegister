package controller

import (
	"example/model"
	"example/service"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ADD MY CODE 0828

// Register User to Postgres
func Register(c *gin.Context) {
	var reqJson model.User
	var err error
	// Receive user struct from request
	if err = c.ShouldBindJSON(&reqJson); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong Req header or json string"})
		return
	}
	// Validate User's structure
	isValid := service.ValidReq(&reqJson)
	if isValid {
		ret, err := model.RegUser(&reqJson)
		if err != nil{
			// Response to FE with 400
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		// Response with 200
		c.IndentedJSON(http.StatusOK, gin.H{"message": ret});
	} else {
		// Response to FE with 400
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Verify your name, email, password"})
	}
}
//[] Register

// Login of User; create token; send only token string to FE.
func Login(c *gin.Context) {
	var user model.User
	// Receive user struct from request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "worng model"})
		return;
	}
	// create token string and store variable result
	result, err := service.LoginUser(&user);
	// err => response with 400 error.
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
	}
	// response token string without JSON.
	c.IndentedJSON(http.StatusOK, gin.H{"message": result} )
}