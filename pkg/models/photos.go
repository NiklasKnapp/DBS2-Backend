package models

import (
	"fmt"
	"mime/multipart"
)

// Photo
type Photo struct {
	Photo_id int     `json:"photoId"`
	Title    string  `json:"title"`
	UUID     string  `json:"uuid" binding:"required"`
	Roll_id  int     `json:"rollId" binding:"required"`
	Rating   float32 `json:"rating"`
}

// Photo upload
type PhotoUpload struct {
	Photo_id int                     `form:"photoid"`
	UUID     string                  `form:"uuid"` //TODO: Change, so UUID is created and added serversided
	Roll_id  int                     `form:"rollId" binding:"required"`
	Files    []*multipart.FileHeader `form:"files" binding:"required"`
}

// Response with photo uploads
type PhotoUploadResponse struct {
	PhotoUpload []PhotoUpload
}

// Create photo in DB
func (p *PhotoUpload) CreatePhoto() (*PhotoUpload, error) {

	// Check for existing film roll
	roll_id, _ := GetFilmRollById(int64(p.Roll_id))
	if roll_id == nil {
		return nil, fmt.Errorf("CreatePhoto: FilmRoll with roll id %v does not exist", p.Roll_id)
	}

	// Run query
	res, err := db.Exec("INSERT INTO photos (title, uuid, roll_id) VALUES (?, ?, ?);", "", p.UUID, p.Roll_id)
	if err != nil {
		return nil, fmt.Errorf("CreatePhoto: %v", err)
	}

	// Get ID
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreatePhoto: %v", err)
	}
	p.Photo_id = int(id)
	return p, nil
}

// Get photos from DB
func GetPhoto() ([]Photo, error) {

	// List of photos
	var photos = []Photo{}

	// Run query
	rows, err := db.Query("SELECT photo_id, title, uuid, roll_id, rating FROM photos;")
	if err != nil {
		return nil, fmt.Errorf("GetPhoto: %v", err)
	}
	defer rows.Close()

	// Extract values
	for rows.Next() {
		var importedPhoto Photo
		if err := rows.Scan(&importedPhoto.Photo_id, &importedPhoto.Title, &importedPhoto.UUID, &importedPhoto.Roll_id, &importedPhoto.Rating); err != nil {
			return nil, fmt.Errorf("GetPhoto: %v", err)
		}
		photos = append(photos, importedPhoto)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPhoto: %v", err)
	}
	return photos, nil
}

// Get photo by type ID from DB
func GetPhotoByTypeId(tId int64) ([]Photo, error) {

	// List of photos
	var photos = []Photo{}

	// Run query
	rows, err := db.Query("SELECT photos.photo_id, photos.title, photos.uuid, photos.roll_id photos.rating FROM photos INNER JOIN film_rolls ON photos.roll_id = film_rolls.roll_id WHERE film_rolls.type_id = ?;", tId)
	if err != nil {
		return nil, fmt.Errorf("GetPhotosByTypeId: %v", err)
	}
	defer rows.Close()

	// Extract values
	for rows.Next() {
		var importetPhotos Photo
		if err := rows.Scan(&importetPhotos.Photo_id, &importetPhotos.Title, &importetPhotos.UUID, &importetPhotos.Roll_id, &importetPhotos.Rating); err != nil {
			return nil, fmt.Errorf("GetPhotosByRollId: %v", err)
		}
		photos = append(photos, importetPhotos)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPhotosByTypeId: %v", err)
	}
	return photos, nil
}

// Get photos by roll ID from DB
func GetPhotosByRollId(rId int64) ([]Photo, error) {

	// List of photos
	var photos = []Photo{}

	// Run query
	rows, err := db.Query("SELECT photo_id, title, uuid, roll_id, rating FROM photos WHERE roll_id = ?;", rId)
	if err != nil {
		return nil, fmt.Errorf("GetPhotosByRollId: %v", err)
	}
	defer rows.Close()

	// Extract values
	for rows.Next() {
		var importetPhotos Photo
		if err := rows.Scan(&importetPhotos.Photo_id, &importetPhotos.Title, &importetPhotos.UUID, &importetPhotos.Roll_id, &importetPhotos.Rating); err != nil {
			return nil, fmt.Errorf("GetPhotosByRollId: %v", err)
		}
		photos = append(photos, importetPhotos)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPhotosByRollId: %v", err)
	}
	return photos, nil
}

// Get photos by album id from DB
func GetPhotosByAlbumId(rId int64) ([]Photo, error) {

	// List of photos
	var photos = []Photo{}

	// Run query
	rows, err := db.Query("SELECT p.photo_id, p.title, p.uuid, p.roll_id, p.rating FROM photos p INNER JOIN photos_albums a ON p.photo_id = a.photo_id WHERE a.album_id = ?;", rId)
	if err != nil {
		return nil, fmt.Errorf("GetPhotosByAlbumId: %v", err)
	}
	defer rows.Close()

	// Extract values
	for rows.Next() {
		var importetPhotos Photo
		if err := rows.Scan(&importetPhotos.Photo_id, &importetPhotos.Title, &importetPhotos.UUID, &importetPhotos.Roll_id, &importetPhotos.Rating); err != nil {
			return nil, fmt.Errorf("GetPhotosByAlbumId: %v", err)
		}
		photos = append(photos, importetPhotos)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPhotosByAlbumId: %v", err)
	}
	return photos, nil
}

// Get photo by ID from DB
func GetPhotoById(pId int64) (*Photo, error) {
	photo := &Photo{}
	if err := db.QueryRow("SELECT photo_id, title, uuid, roll_id, rating FROM photos WHERE photo_id = ?;", pId).Scan(&photo.Photo_id, &photo.Title, &photo.UUID, &photo.Roll_id, &photo.Rating); err != nil {
		return nil, fmt.Errorf("GetPhotoById: %v", err)
	}
	return photo, nil
}

// Update photo in DB
func (p *Photo) UpdatePhoto() (*Photo, error) {
	_, err := db.Exec("UPDATE photos SET title = ?, uuid = ?, roll_id = ?, rating = ? WHERE photo_id = ?;", p.Title, p.UUID, p.Roll_id, p.Rating, p.Photo_id)
	if err != nil {
		return nil, fmt.Errorf("UpdatePhoto: %v", err)
	}
	return p, nil
}

//Delete photo in DB
func DeletePhoto(pId int64) (*Photo, error) {
	photo, _ := GetPhotoById(pId)
	_, err := db.Exec("DELETE FROM photos WHERE photo_id = ?;", pId)
	if err != nil {
		return nil, fmt.Errorf("DeletePhoto: %v", err)
	}
	return photo, nil
}
