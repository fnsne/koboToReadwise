package models

type Content struct {
	ContentID string `gorm:"primaryKey;column:ContentID"`
	Title     string `gorm:"column:Title"`
	BookTitle string
}

func (c *Content) TableName() string {
	return "content"
}

func (c *Content) Author() string {
	return "From koboToReadwise Tool"
}
