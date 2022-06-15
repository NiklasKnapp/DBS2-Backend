package models

import (
	"fmt"
)

type Rating struct {
	Rating_id   int `json:"ratingId"`
	Photo_id    int `json:"photoId" binding:"required"`
	Rating      int `json:"rating"`
}

func (r *Rating) CreateRating() (*Rating, error) {
	photo_id, _ := GetPhotoById(int64(r.Photo_id))
	if photo_id == nil {
		return nil, fmt.Errorf("CreateRating: Photo with photo id %v does not exist", r.Photo_id)
	}
	res, err := db.Exec("INSERT INTO ratings (photo_id, rating) VALUES (?, ?);", "", r.Photo_id, r.Rating)
	if err != nil {
		return nil, fmt.Errorf("CreateRating: %v", err)
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateRating: %v", err)
	}
	r.Rating_id = int(id)
	return r, nil
}

func GetRating() ([]Rating, error) {
	var Ratings = []Rating{}
	rows, err := db.Query("SELECT rating_id, photo_id, rating FROM ratings;")
	if err != nil {
		return nil, fmt.Errorf("GetRating: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var ratings Rating
		if err := rows.Scan(&ratings.Rating_id, &ratings.Photo_id, &ratings.Rating); err != nil {
			return nil, fmt.Errorf("GetRating: %v", err)
		}
		Ratings = append(Ratings, ratings)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetRating: %v", err)
	}
	return Ratings, nil
}

func GetRatingById(rId int64) (*Rating, error) {
	rating := &Rating{}
	if err := db.QueryRow("SELECT rating_id, photo_id, rating FROM ratings WHERE rating_id = ?;", rId).Scan(&rating.Rating_id, &rating.Photo_id, &rating.Rating); err != nil {
		return nil, fmt.Errorf("GetRatingById: %v", err)
	}
	return rating, nil
}

func (fr *Rating) UpdateRating() (*Rating, error) {
	_, err := db.Exec("UPDATE ratings SET rating = ?, photoId = ? WHERE rating_id = ?;", fr.Rating, fr.Photo_id, fr.Rating_id)
	if err != nil {
		return nil, fmt.Errorf("UpdateRating: %v", err)
	}
	return fr, nil
}

func DeleteRating(rId int64) (*Rating, error) {
	rating, _ := GetRatingById(rId)
	_, err := db.Exec("DELETE FROM ratings WHERE rating_id = ?;", rId)
	if err != nil {
		return nil, fmt.Errorf("DeleteRating: %v", err)
	}
	return rating, nil
}
