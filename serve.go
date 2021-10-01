package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teng231/workshop/db"
)

type Workshop struct {
	db *db.DB
}

func startServer() {

	router := gin.Default()
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	// router.GET("/welcome", func(c *gin.Context) {
	// 	firstname := c.DefaultQuery("firstname", "Guest")
	// 	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	// })
	router.GET("/users", handleListUsers)
	router.POST("/users", handleCreateUser)
	router.POST("/users/:id", handleUpdateUserById)
	router.GET("/users/:id", handleGetUserById)
	router.Run(":8080")
}
