package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

  "crud_restapi/handler"
)


func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/books", handlers.ListBooks(db))
	r.GET("/books/:id", handlers.GetBook(db))
	r.POST("/books", handlers.CreateBook(db))
	r.PUT("/books/:id", handlers.UpdateBook(db))
	r.DELETE("/books/:id", handlers.DeleteBook(db))
}
