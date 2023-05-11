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

const (
	mac     = "mac"
	windows = "windows"
)

var OS string

func main() {
	env, err := GetEnvs(OS)
	if err != nil {
		panic(fmt.Errorf("cannot get envs, error=%v", err))
	}

	var sqlPosition string
	switch OS {
	case mac:
		sqlPosition = path.Join(env.Homedir, "Library/Application Support/", "Kobo", "Kobo Desktop Edition", "Kobo.sqlite")
	case windows:
		sqlPosition = path.Join(env.Homedir, "AppData", "Local", "Kobo", "Kobo Desktop Edition", "Kobo.sqlite")
	default:
		panic(fmt.Errorf("wrong os, os=%v", OS))
	}

	fmt.Println("kobo.sqlite position = ", sqlPosition)
	db, err := gorm.Open(sqlite.Open(sqlPosition), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot open gorm, error=%v", err))
	}
	puller := repos.NewBookmarkPuller(db)
	bookmarks, err := puller.GetBookmarkList()

	err = WriteKindleClippingFormat("output.txt", bookmarks)
	if err != nil {
		panic(fmt.Errorf("cannot write kindle clipping format, error=%v", err))
	}
}

type Env struct {
	Homedir string
}

func GetEnvs(osCode string) (Env, error) {
	var homeDir string
	switch osCode {
	case mac:
		homeDir = os.Getenv("HOME")
	case windows:
		homeDir = os.Getenv("USERPROFILE")
	}
	env := Env{
		Homedir: homeDir,
	}
	return env, nil
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
