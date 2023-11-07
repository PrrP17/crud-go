package routes

import (
	"github.com/PrrP17/crud-go.git/src/contreoller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", contreoller.FindUserByID)

	r.GET("/getUserByEmail/:userEmail", contreoller.FindUserByEmail)

	r.POST("/createUser", contreoller.CreateUser)

	r.PUT("/updateUser/:userId", contreoller.UpdateUser)

	r.DELETE("/deleteUser/:userId", contreoller.DeleteUser)
}
