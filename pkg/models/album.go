package models

import (
	"fmt"
)

type Album struct {
	Album_id     int     `json:"albumId"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Uuid        string  `json:"uuid"`
	Imagerating float32 `json:"imageRating"`
}

type PA struct {
	Album_id   int  `json:"albumId"`
	Photo_id   int  `json:"photoId`
}

func GetAlbum() ([]Album, error) {
	var albumse = []Album{}
	rows, err := db.Query("SELECT albums.album_id, albums.title, albums.text, albums.rating FROM albums;")
	if err != nil {
		return nil, fmt.Errorf("GetAlbum: %v", err)
	}
	
	defer rows.Close()
	for rows.Next() {
		var albums Album
		if err := rows.Scan(&albums.Album_id, &albums.Title, &albums.Description, &albums.Rating); err != nil {
			return nil, fmt.Errorf("GetAlbum: %v", err)
		}
		albums.Uuid = ""
		albums.Imagerating = 0

		row, err := db.Query("SELECT albums.album_id, albums.title, albums.text, albums.rating, photos.uuid, MAX(photos.rating) FROM albums LEFT JOIN photos_albums ON albums.album_id = photos_albums.album_id LEFT JOIN photos ON photos_albums.photo_id = photos.photo_id GROUP BY albums.album_id HAVING albums.album_id = ?;", albums.Album_id)
		if err != nil {
			return nil, fmt.Errorf("GetAlbum: %v", err)
		}
		var album Album
		row.Next()
		if err := row.Scan(&album.Album_id, &album.Title, &album.Description, &album.Rating, &album.Uuid, &album.Imagerating); err != nil {
			albumse = append(albumse, albums)
		} else {
			albumse = append(albumse, album)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAlbum: %v", err)
	}
	return albumse, nil
}

func GetAlbumById(rId int64) (*Album, error) {
	album := &Album{}
	if err := db.QueryRow("SELECT album_id, title, text, type_id, rating FROM albums WHERE album_id = ?;", rId).Scan(&album.Album_id, &album.Title, &album.Description, &album.Rating); err != nil {
		return nil, fmt.Errorf("GetAlbumById: %v", err)
	}
	return album, nil
}

func (fr *Album) UpdateAlbum() (*Album, error) {
	_, err := db.Exec("UPDATE albums SET title = ?, text = ?, rating = ? WHERE album_id = ?;", fr.Title, fr.Description, fr.Album_id, fr.Rating)
	if err != nil {
		return nil, fmt.Errorf("UpdateAlbum: %v", err)
	}
	return fr, nil
}

func DeleteAlbum(rId int64) (*Album, error) {
	album, _ := GetAlbumById(rId)
	_, err := db.Exec("DELETE FROM albums WHERE album_id = ?;", rId)
	if err != nil {
		return nil, fmt.Errorf("DeleteAlbum: %v", err)
	}
	return album, nil
}

func (fr *PA) CreatePA() (*PA, error) {
	photo_id, _ := GetPhotoById(int64(fr.Photo_id))
	if photo_id == nil {
		return nil, fmt.Errorf("CreatePA: Photo with photo_id %v does not exist", fr.Photo_id)
	}
	album_id, _ := GetAlbumById(int64(fr.Album_id))
	if album_id == nil {
		return nil, fmt.Errorf("CreatePA: Album with album_id %v does not exist", fr.Album_id)
	}
	_, err := db.Exec("INSERT INTO photos_albums (photo_id, album_id) VALUES(?, ?);", fr.Photo_id, fr.Album_id)
	if err != nil {
		return nil, fmt.Errorf("CreatePA: %v", err)
	}
	return fr, nil
}