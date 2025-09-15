package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
  "crud_restapi/models"
)


func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		if limit <= 0 || limit > 100 {
			limit = 20
		}
		offset := (page - 1) * limit
		q := c.Query("q")

		var books []models.Book
		tx := db.Model(&models.Book{})
		if q != "" {
			like := "%" + q + "%"
			tx = tx.Where("title LIKE ? OR author LIKE ?", like, like)
		}

		var total int64
		tx.Count(&total)

		if err := tx.Order("id DESC").Limit(limit).Offset(offset).Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch books"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  books,
			"page":  page,
			"limit": limit,
			"total": total,
		})
	}
}

func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Book
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		if err := db.Create(&in).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create"})
			return
		}
		c.JSON(http.StatusCreated, in)
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}

		var in models.Book
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}

		book.Title = in.Title
		book.Author = in.Author
		book.Price = in.Price

		if err := db.Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Book{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}