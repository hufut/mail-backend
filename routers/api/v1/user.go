package v1

import (
	"github.com/gin-gonic/gin"

	_ "github.com/hukaixuan/mall-backend/pkg/e"
)

// @Summary Get user detail
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} string 200
// @Failure 500
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {

}

func Registe(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func EditUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
