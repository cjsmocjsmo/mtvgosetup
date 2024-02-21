package mtvgosetup

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	// "github.com/disintegration/imaging"
)

func MtvGoSetup() {
	dbdirpath := os.Getenv("MTV_DB_DIR_PATH")
	if _, err := os.Stat(dbdirpath); os.IsNotExist(err) {
		os.Mkdir(dbdirpath, 0755)
	}

	dbpath := os.Getenv("MTV_DB_PATH")
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		os.Create(dbpath)
	}

	CreateImagesDB(dbpath)
	CreateMoviesDB(dbpath)
	CreateTVShowsDB(dbpath)

	CreateThumbnails()
}

func CreateThumbnails()  {
	poster_path := os.Getenv("MTV_POSTER_PATH")
	// thumb_list := []string{}
	err := filepath.Walk(poster_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		thumbnailPath := os.Getenv("MTV_THUMBNAILS_PATH") // Replace with the actual path where you want to save the thumbnail
		fmt.Println(thumbnailPath)
		fmt.Println(path)
	
		// 	img, err := imaging.Open(path)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	thumbnail := imaging.Resize(img, 800, 0, imaging.Lanczos)
	// 	err = imaging.Save(thumbnail, thumbnailPath)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	thumb_list = append(thumb_list, thumbnailPath)

		return nil
	})
	if err != nil {
		panic(err)
	}
	// return thumb_list
}

func CreateMoviesDB(db_path string) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS movies (
		id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            year TEXT NOT NULL,
            posteraddr TEXT NOT NULL,
            size TEXT NOT NULL,
            path TEXT NOT NULL,
            idx TEXT NOT NULL,
            movid TEXT NOT NULL UNIQUE,
            catagory TEXT NOT NULL,
            httpthumbpath TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func CreateTVShowsDB(db_path string) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tvshows (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
            tvid TEXT NOT NULL UNIQUE,
            size TEXT NOT NULL,
            catagory TEXT NOT NULL,
            name TEXT NOT NULL,
            season TEXT NOT NULL,
            episode TEXT NOT NULL,
            path TEXT NOT NULL,
            idx TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func CreateImagesDB(db_path string) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
            imgid TEXT NOT NULL UNIQUE,
            path TEXT NOT NULL,
            imgpath TEXT NOT NULL,
            size TEXT NOT NULL,
            name TEXT NOT NULL,
            thumbpath TEXT NOT NULL,
            idx INTEGER NOT NULL,
            httpthumbpath TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
