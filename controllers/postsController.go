package controllers

import (
	"github.com/apanji404/Prakerja_Golang/initializers"
	"github.com/apanji404/Prakerja_Golang/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var val struct {
		Menu  string
		Price int
	}

	c.Bind(&val)

	post := models.Post{Menu: val.Menu, Price: val.Price}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsGet(c *gin.Context) {
	var getDB []models.Post
	initializers.DB.Find(&getDB)

	c.JSON(200, gin.H{
		"getDB": getDB,
	})
}

func PostsGetByID(c *gin.Context) {
	id := c.Param("id")

	var findID models.Post
	result := initializers.DB.First(&findID, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"findDB": findID,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var updateDB struct {
		Menu  string
		Price int
	}

	c.Bind(&updateDB)

	var findID models.Post
	result := initializers.DB.First(&findID, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}

	initializers.DB.Model(&findID).Updates(models.Post{
		Menu:  updateDB.Menu,
		Price: updateDB.Price,
	})

	c.JSON(200, gin.H{
		"updateDB": updateDB,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}

	c.Status(200)
}
