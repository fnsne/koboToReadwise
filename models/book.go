package models

type Book struct {
	ContentID     string `gorm:"primaryKey;column:ContentID"`
	BookTitle     string `gorm:"column:Title"`
	Accessibility int    `gorm:"column:Accessibility"`
	Author        string `gorm:"column:Attribution"`
	//SubTitle     string
	//Author       string
	//Publisher    string
	//ISBN         string
	//ReleaseDate  string
	//Series       string
	//SeriesNumber int
	//Rating       int
	//ReadPercent  int
	//LastRead     string
	//FileSize     int
	//Source       string
}

func (b *Book) TableName() string {
	return "content"
}
