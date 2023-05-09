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

	for i := 0; i < len(bookmarks); i++ {
		bookmark := bookmarks[i]
		book, err := b.GetBookByContentID(bookmark.ContentID)
		if err != nil {
			return nil, fmt.Errorf("cannot get book by content id, error=%v", err)
		}
		bookmarks[i].Book = book
	}

	sort.Slice(bookmarks, func(i, j int) bool {
		return bookmarks[i].Content.Title > bookmarks[j].Content.Title
	})
	return bookmarks, err
}

func (b *BookmarkPuller) GetBookList() ([]models.Book, error) {
	var books []models.Book
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

func (b *BookmarkPuller) GetBookByContentID(contentID string) (models.Book, error) {
	actualContentID := strings.Split(contentID, "!")[0]
	var book models.Book
	err := b.db.Where("ContentID=?", actualContentID).Find(&book).Error
	return book, err
}
