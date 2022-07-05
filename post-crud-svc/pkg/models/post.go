package models

type Post struct {
	ID     int64 `gorm:"primarykey,autoIncrement:false"`
	UserID int64
	Title  string
	Body   string
}
