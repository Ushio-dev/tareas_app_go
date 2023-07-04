package controllers

import (
	"github.com/Ushio-dev/tareas-app/initializers"
	"github.com/Ushio-dev/tareas-app/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return data
	c.JSON(200, gin.H{
		"msg": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var post []models.Post
	initializers.DB.Find(&post)

	// Response
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostShow(c *gin.Context) {
	// Get Id off url
	id := c.Param("id")
	//get the post
	var post models.Post
	initializers.DB.First(&post, id)
	//respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get de id
	id := c.Param("id")
	// delete de post
	initializers.DB.Delete(&models.Post{}, id)
	// respond
	c.Status(200)
}
