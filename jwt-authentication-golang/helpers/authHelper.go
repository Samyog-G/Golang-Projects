package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized access")
		return err
	}
	return err
}
func MatchUserTypeToUid(c *gin.Context, userID string) (err error) {
	userType := c.GetString("user type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userID {
		err = errors.New("Unauthorized access")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}