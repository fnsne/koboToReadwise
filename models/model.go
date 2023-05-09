package models

import (
	"fmt"
	"math"
	"time"
)

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

type Bookmark struct {
	BookmarkID      string `gorm:"primaryKey"`
	Text            string `gorm:"column:text"`
	ChapterProgress float64
	Annotation      string
	ContentID       string
	Content         Content
	DateCreated     string `gorm:"column:DateCreated"`
}

func (b *Bookmark) CreateTime() time.Time {
	t, err := time.Parse("2006-01-02T15:04:05.000", b.DateCreated)
	if err != nil {
		fmt.Printf("Parse time <%s> got error=%v\n", b.DateCreated, err)
	}
	return t
}

func (b *Bookmark) Author() string {
	return b.Content.Author()
}

func (b *Bookmark) TableName() string {
	return "Bookmark"
}

func (b *Bookmark) Location() float64 {
	location := math.Round(b.ChapterProgress * 100)
	return location
}

func (b *Bookmark) HighLight() string {
	return b.Text
}

func (b *Bookmark) Output() string {
	var str string
	str += fmt.Sprintf("%s (%s)\n", b.Content.BookTitle, b.Author())
	str += fmt.Sprintf("- Your Highlight on Location %.0f | Added on %s", b.Location(), b.CreateTime().Format("Monday, January 2, 2006 15:04:05 PM"))
	str += fmt.Sprintf("\n\n%s", b.HighLight())
	str += fmt.Sprintf("\n==========")

	if b.Annotation != "" {
		str += fmt.Sprintf("\n%s (%s)\n", b.Content.BookTitle, b.Author())
		str += fmt.Sprintf("- Your Note on Location %.0f | Added on %s", b.Location(), b.CreateTime().Format("Monday, January 2, 2006 15:04:05 PM"))
		str += fmt.Sprintf("\n\n%s", b.Annotation)
		str += fmt.Sprintf("\n==========")
	}

	return str
}
