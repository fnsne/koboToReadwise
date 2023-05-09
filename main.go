package main

import (
	"bufio"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"koboToReadwise/models"
	"koboToReadwise/repos"
	"os"
	"path"
)

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

	fmt.Println("kobo.sqlite position = ", sqlPosition)
	db, err := gorm.Open(sqlite.Open(sqlPosition), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot open gorm, error=%v", err))
	}
	puller := repos.NewBookmarkPuller(db)
	//bookList, err := puller.GetBookList()
	//if err != nil {
	//	panic(fmt.Errorf("cannot get book list, error=%v", err))
	//}
	//for _, book := range bookList {
	//	fmt.Printf("book title = %v\n", book.BookTitle)
	//	fmt.Printf("author = %v\n", book.Author)
	//	fmt.Printf("content id = %v\n", book.ContentID)
	//	fmt.Printf("accessibility = %v\n", book.Accessibility)
	//}

	bookmarks, err := puller.GetBookmarkList()

	err = WriteKindleClippingFormat("output.txt", bookmarks)
	if err != nil {
		panic(fmt.Errorf("cannot write kindle clipping format, error=%v", err))
	}
}

func WriteKindleClippingFormat(outputFileName string, bookmarks []models.Bookmark) error {
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("cannot create file, error=%v", err)
	}
	defer file.Close()

	// 將字串切片寫入檔案
	writer := bufio.NewWriter(file)
	for _, bookmark := range bookmarks {
		_, err := writer.WriteString(bookmark.Output() + "\n")
		if err != nil {
			return fmt.Errorf("cannot write string, error=%v", err)
		}
	}
	return writer.Flush()
}
