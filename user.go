package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teng231/workshop/pb"
)

func handleListUsers(c *gin.Context) {
	list, err := ws.db.ListUsers(&pb.UserRequest{})
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, list)
}

func handleCreateUser(c *gin.Context) {
	user := &pb.User{}
	c.ShouldBindJSON(user)
	err := ws.db.InsertUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"error": "success"})
}

func handleUpdateUserById(c *gin.Context) {
	user := &pb.User{}
	c.ShouldBindJSON(user)
	if user.GetId() == 0 {
		c.JSON(400, gin.H{"error": "not found id"})
		return
	}
	err := ws.db.UpdateUser(user, &pb.User{Id: user.Id})
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"error": "success"})
}

func handleGetUserById(c *gin.Context) {
	if c.Query("id") == "" {
		c.JSON(400, gin.H{"error": "not found id"})
		return
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	user, err := ws.db.FindUser(&pb.User{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, user)
}
