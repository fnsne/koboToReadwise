package models

type Book struct {
	ContentID     string `gorm:"primaryKey;column:ContentID"`
	BookTitle     string `gorm:"column:Title"`
	Accessibility int    `gorm:"column:Accessibility"`
	Author        string `gorm:"column:Attribution"`
}

func (b *Book) TableName() string {
	return "content"
}
