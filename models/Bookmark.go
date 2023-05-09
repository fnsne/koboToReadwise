package models

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Bookmark struct {
	BookmarkID      string `gorm:"primaryKey"`
	Text            string `gorm:"column:text"`
	ChapterProgress float64
	Annotation      string
	ContentID       string
	DateCreated     string `gorm:"column:DateCreated"`
	Book            Book   `gorm:"-"`
}

func (b *Bookmark) Author() string {
	return strings.Split(b.Book.Author, ",")[0]
}

func (b *Bookmark) TableName() string {
	return "Bookmark"
}

func (b *Bookmark) Location() float64 {
	location := math.Round(b.ChapterProgress * 100)
	return location
}

func (b *Bookmark) CreateTime() time.Time {
	t, err := time.Parse("2006-01-02T15:04:05.000", b.DateCreated)
	if err != nil {
		fmt.Printf("Parse time <%s> got error=%v\n", b.DateCreated, err)
	}
	return t
}

func (b *Bookmark) Output() string {
	var str string
	str += fmt.Sprintf("%s (%s)\n", b.BookTitle(), b.Author())
	str += fmt.Sprintf("- Your Highlight on Location %.0f | Added on %s", b.Location(), b.CreateTime().Format("Monday, January 2, 2006 15:04:05 PM"))
	str += fmt.Sprintf("\n\n%s", b.Text)
	str += fmt.Sprintf("\n==========")

	if b.hasNote() {
		str += fmt.Sprintf("\n%s (%s)\n", b.BookTitle(), b.Author())
		str += fmt.Sprintf("- Your Note on Location %.0f | Added on %s", b.Location(), b.CreateTime().Format("Monday, January 2, 2006 15:04:05 PM"))
		str += fmt.Sprintf("\n\n%s", b.Annotation)
		str += fmt.Sprintf("\n==========")
	}

	return str
}

func (b *Bookmark) hasNote() bool {
	hasNote := b.Annotation != ""
	return hasNote
}

func (b *Bookmark) BookTitle() string {
	return b.Book.BookTitle
}
