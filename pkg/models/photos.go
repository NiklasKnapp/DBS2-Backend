package models

import "fmt"

type Photo struct {
	Photo_id int    `json:"photoId"`
	Title    string `json:"title"`
	UUID     string `json:"uuid" binding:"required"`
	Roll_id  int    `json:"rollId" binding:"required"`
}

func (p *Photo) CreatePhoto() (*Photo, error) {
	roll_id, _ := GetFilmRollById(int64(p.Roll_id))
	if roll_id == nil {
		return nil, fmt.Errorf("CreatePhoto: FilmRoll with roll id %v does not exist", p.Roll_id)
	}
	res, err := db.Exec("INSERT INTO photos (title, uuid, roll_id) VALUES (?, ?, ?);", p.Title, p.UUID, p.Roll_id)
	if err != nil {
		return nil, fmt.Errorf("CreatePhoto: %v", err)
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreatePhoto: %v", err)
	}
	p.Photo_id = int(id)
	return p, nil
}

func GetPhoto() ([]Photo, error) {
	var photos = []Photo{}
	rows, err := db.Query("SELECT photo_id, title, uuid, roll_id FROM photos;")
	if err != nil {
		return nil, fmt.Errorf("GetPhoto: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var importedPhoto Photo
		if err := rows.Scan(&importedPhoto.Photo_id, &importedPhoto.Title, &importedPhoto.UUID, &importedPhoto.Roll_id); err != nil {
			return nil, fmt.Errorf("GetPhoto: %v", err)
		}
		photos = append(photos, importedPhoto)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPhoto: %v", err)
	}
	return photos, nil
}

func GetPhotoById(pId int64) (*Photo, error) {
	photo := &Photo{}
	if err := db.QueryRow("SELECT photo_id, title, uuid, roll_id FROM photos WHERE photo_id = ?;", pId).Scan(&photo.Photo_id, &photo.Title, &photo.UUID, &photo.Roll_id); err != nil {
		return nil, fmt.Errorf("GetPhotoById: %v", err)
	}
	return photo, nil
}

func (p *Photo) UpdatePhoto() (*Photo, error) {
	_, err := db.Exec("UPDATE photos SET title = ?, uuid = ?, roll_id = ? WHERE photo_id = ?;", p.Title, p.UUID, p.Roll_id, p.Photo_id)
	if err != nil {
		return nil, fmt.Errorf("UpdatePhoto: %v", err)
	}
	return p, nil
}

func DeletePhoto(pId int64) (*Photo, error) {
	photo, _ := GetPhotoById(pId)
	_, err := db.Exec("DELETE FROM photos WHERE photo_id = ?;", pId)
	if err != nil {
		return nil, fmt.Errorf("DeletePhoto: %v", err)
	}
	return photo, nil
}
