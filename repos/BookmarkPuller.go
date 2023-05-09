package repos

import (
	"fmt"
	"gorm.io/gorm"
	"koboToReadwise/models"
	"sort"
	"strings"
)

type BookmarkPuller struct {
	db *gorm.DB
}

func NewBookmarkPuller(db *gorm.DB) BookmarkPuller {
	return BookmarkPuller{db: db}
}

func (b *BookmarkPuller) GetBookmarkList() ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	err := b.db.Preload("Content").Find(&bookmarks).Error
	if err != nil {
		panic(fmt.Errorf("cannot get bookmarks, error=%v", err))
	}
	sort.Slice(bookmarks, func(i, j int) bool {
		return bookmarks[i].Content.Title > bookmarks[j].Content.Title
	})
	return bookmarks, err
}

type Book struct {
	ContentID     string `gorm:"primaryKey;column:ContentID"`
	BookTitle     string `gorm:"column:Title"`
	IsPurchasable bool   `gorm:"column:IsPurchasable"`
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

func (b *BookmarkPuller) GetBookList() ([]Book, error) {
	var books []Book
	err := b.db.
		Where("ContentType=?", 6).
		Where("___UserID is not null").
		Where("___UserID !=''").
		Where("___UserID !=?", "removed").
		Where("accessibility=?", 1). //1: store購買的, 6: Preview
		Find(&books).Error
	for i := 0; i < len(books); i++ {
		book := books[i]
		books[i].Author = strings.Split(book.Author, ",")[0]
	}
	return books, err
}
