package v1

import (
	"github.com/gin-gonic/gin"

	"admin/model/response"
)

func GetUserInfo(c *gin.Context) {
	response.SuccessResponse(c, "user")
}
