package controllers

import (
	"github.com/19abhishek/todo/initializers"
	"github.com/19abhishek/todo/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context){
	//Get data off req body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context){
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context){
	//Get id from the URL
	id := c.Param("id")

	//Get the Post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context){
	// Get the id of the URL
	id := c.Param("id")

	// Get the data of the request body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Find the Post you're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context){
	//Get the id from the URL
	id := c.Param("id")

	//Delete the Posts
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}