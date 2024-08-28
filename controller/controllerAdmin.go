package controller

import (
	"example/model"
	// "example/service"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ADD MY CODE 0828
// Get users array from postgres
func GetUsers(c *gin.Context) {
	// if parameter exists, mUuid is parameter value, else, mUuid is "" empty string.
	// mUuid := c.Param("user_id")
	mUuid := c.Query("user_id")
	// get user info
	ret, err := model.GetUsers(mUuid)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err});
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": ret});
}

func DelUserUID(c *gin.Context) {
	// mUuid equals to parameter value.
	// mUuid := c.Param("user_id")
	mUuid := c.Query("user_id")
	// get Data from model based on postgres.
	res := model.DelUserUID(mUuid)

	if res {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Del usersUID by admin"})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Del usersUID by admin"})
	return
}

func UsersUID(c *gin.Context) {
	// mUuid equals to parameter value.
	// mUuid := c.Param("user_id")
	mUuid := c.Query("user_id")
	// get user's structure from request.
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong model"})
		return;
	}

	// update user information in UpdateUser model by ADMIN.
	if res := model.UpdateUserAdmin(mUuid, &user) && mUuid != ""; res {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "post update usersUID by admin"})
		return
	}
	// update user information in UpdateUser model by USER.
	if res := model.UpdateUser(&user) && mUuid == ""; res {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "post update usersUID by admin"})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No match condition"})
}
