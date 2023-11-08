package contreoller

import (
	"fmt"
	"github.com/PrrP17/crud-go.git/src/configuration/rest_err"
	"github.com/PrrP17/crud-go.git/src/contreoller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestErr(
			fmt.Sprintf("There are some incorrect fildes, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
}
