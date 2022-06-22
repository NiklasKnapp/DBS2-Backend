package models

import (
	"fmt"
	"strconv"
)

// Album
type Album struct {
	Album_id    int     `json:"albumId"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Uuid        string  `json:"uuid"`
	Imagerating float32 `json:"imageRating"`
}

// Link photos_albums
type PA struct {
	Album_id   string  `json:"albumId"`
	Photo_id   string  `json:"photoId"`
}

//Create album in DB
func (fr *Album) CreateAlbum() (*Album, error) {

	// Run query
	res, err := db.Exec("INSERT INTO albums (title, text) VALUES(?, ?);", fr.Title, fr.Description)
	if err != nil {
		return nil, fmt.Errorf("CreateAlbum: %v", err)
	}

	// Get ID
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateAlbum: %v", err)
	}
	fr.Album_id = int(id)
	fr.Rating = 0
	return fr, nil
}

// Get albums from DB
func GetAlbum() ([]Album, error) {

	// List of albums
	var albumse = []Album{}

	// Run query
	rows, err := db.Query("SELECT albums.album_id, albums.title, albums.text, albums.rating FROM albums;")
	if err != nil {
		return nil, fmt.Errorf("GetAlbum: %v", err)
	}
	defer rows.Close()

	// Extract values
	for rows.Next() {
		var albums Album
		if err := rows.Scan(&albums.Album_id, &albums.Title, &albums.Description, &albums.Rating); err != nil {
			return nil, fmt.Errorf("GetAlbum: %v", err)
		}

		// Get rating and best rated photo in album
		albums.Uuid = ""
		albums.Imagerating = 0

		row, err := db.Query("SELECT albums.album_id, albums.title, albums.text, albums.rating, photos.uuid, MAX(photos.rating) FROM albums LEFT JOIN photos_albums ON albums.album_id = photos_albums.album_id LEFT JOIN photos ON photos_albums.photo_id = photos.photo_id GROUP BY albums.album_id HAVING albums.album_id = ?;", albums.Album_id)
		if err != nil {
			return nil, fmt.Errorf("GetAlbum: %v", err)
		}
		var album Album
		row.Next()

		// Album with images
		if err := row.Scan(&album.Album_id, &album.Title, &album.Description, &album.Rating, &album.Uuid, &album.Imagerating); err != nil {
			albumse = append(albumse, albums)
		
		// Empty album
		} else {
			albumse = append(albumse, album)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAlbum: %v", err)
	}
	return albumse, nil
}

// Get album by ID from DB
func GetAlbumById2(rId int64) (*Album, error) {
	album := &Album{}
	if err := db.QueryRow("SELECT album_id, title, text, rating FROM albums WHERE album_id = ?;", rId).Scan(&album.Album_id, &album.Title, &album.Description, &album.Rating); err != nil {
		return nil, fmt.Errorf("GetAlbumById: %v", err)
	}
	return album, nil
}

// Update album in DB
func (fr *Album) UpdateAlbum() (*Album, error) {
	_, err := db.Exec("UPDATE albums SET title = ?, text = ?, rating = ? WHERE album_id = ?;", fr.Title, fr.Description, fr.Album_id, fr.Rating)
	if err != nil {
		return nil, fmt.Errorf("UpdateAlbum: %v", err)
	}
	return fr, nil
}

// Delete album from DB
func DeleteAlbum(rId int64) (*Album, error) {
	album, _ := GetAlbumById2(rId)
	_, err := db.Exec("DELETE FROM albums WHERE album_id = ?;", rId)
	if err != nil {
		return nil, fmt.Errorf("DeleteAlbum: %v", err)
	}
	return album, nil
}

//Create Link photo album in DB
func (fr *PA) CreatePA() (*PA, error) {

	// Check for existing photo
	p, _ :=strconv.Atoi(fr.Photo_id)
	photo_id, _ := GetPhotoById(int64(p))
	if photo_id == nil {
		return nil, fmt.Errorf("CreatePA: Photo with photo_id %v does not exist", fr.Photo_id)
	}

	// Check for existing album
	a, _ := strconv.Atoi(fr.Album_id)
	album_id, _ := GetAlbumById2(int64(a))
	if album_id == nil {
		return nil, fmt.Errorf("CreatePA: Album with album_id %v does not exist", fr.Album_id)
	}

	// Run query
	_, err := db.Exec("INSERT INTO photos_albums (photo_id, album_id) VALUES(?, ?);", fr.Photo_id, fr.Album_id)
	if err != nil {
		return nil, fmt.Errorf("CreatePA: %v", err)
	}
	return fr, nil
}