package utils

import (
	"BusinessSys/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

func GetUser(c *gin.Context) (user *model.User,err error) {
	var (
		data interface{}
		getOK bool
		transformOK bool
	)

	if data,getOK = c.Get("currentUser");!getOK {
		return nil,errors.New("currentUser is not exist")
	}

	if user,transformOK = data.(*model.User);!transformOK{
		return nil,errors.New("session data transform fail")
	}

	return
}
