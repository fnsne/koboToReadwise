package main

import (
	"bufio"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math"
	"os"
	"path"
	"sort"
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

func main() {
	var isMac, isWindows bool
	//todo change to use flag or env
	isMac = true
	isWindows = false
	var sqlPosition string
	if isMac {
		sqlPosition = path.Join("/Users/fnsne/Library/Application Support/", "Kobo", "Kobo Desktop Edition", "Kobo.sqlite")
	} else if isWindows {
		sqlPosition = path.Join("C://Users/watas/", "AppData", "Local", "Kobo", "Kobo Desktop Edition", "Kobo.sqlite")
	}

	fmt.Println("sqlPosition=", sqlPosition)
	db, err := gorm.Open(sqlite.Open(sqlPosition), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot open gorm, error=%v", err))
	}
	var bookmarks []Bookmark
	err = db.Preload("Content").Find(&bookmarks).Error
	if err != nil {
		panic(fmt.Errorf("cannot get bookmarks, error=%v", err))
	}
	sort.Slice(bookmarks, func(i, j int) bool {
		return bookmarks[i].Content.Title > bookmarks[j].Content.Title
	})

	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 將字串切片寫入檔案
	writer := bufio.NewWriter(file)
	for _, bookmark := range bookmarks {
		_, err := writer.WriteString(bookmark.Output() + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
