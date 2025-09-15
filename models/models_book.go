package models

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" binding:"required" gorm:"type:varchar(255);index"`
	Author string `json:"author" binding:"required" gorm:"type:varchar(255);index"`
	Price  int    `json:"price"`
}
