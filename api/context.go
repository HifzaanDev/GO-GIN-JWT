package main

import (
	"rest-api-in-gin/internal/database"

	"github.com/gin-gonic/gin"
)

func (app *application) GetUserFromContext(c *gin.Context) *database.User {
	contextUser, exist := c.Get("user")
	if !exist {
		return nil
	}
	user, ok := contextUser.(*database.User)
	if !ok {
		return nil
	}

	return user
}
