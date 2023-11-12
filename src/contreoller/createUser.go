package contreoller

import (
	"github.com/PrrP17/crud-go.git/src/configuration/logger"
	"github.com/PrrP17/crud-go.git/src/configuration/validation"
	"github.com/PrrP17/crud-go.git/src/contreoller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		errRest := validation.ValidationError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
}
