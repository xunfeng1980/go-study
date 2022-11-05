package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// load the casbin model and policy from files, database is also supported.
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
	router := gin.New()
	router.Use(authz.NewAuthorizer(e))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
