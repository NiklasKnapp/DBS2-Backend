package models

import (
	"fmt"
)

type Rating struct {
	Rating_id   int `json:"ratingId"`
	Photo_id    int `json:"photoId" binding:"required"`
	rating      int `json:"rating"`
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
		if err := rows.Scan(&ratings.Rating_id, &ratings.Photo_id, &ratings.rating); err != nil {
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
	if err := db.QueryRow("SELECT rating_id, photo_id, rating FROM ratings WHERE rating_id = ?;", rId).Scan(&rating.Rating_id, &rating.Photo_id, &rating.rating); err != nil {
		return nil, fmt.Errorf("GetRatingById: %v", err)
	}
	return rating, nil
}

func (fr *Rating) UpdateRating() (*Rating, error) {
	_, err := db.Exec("UPDATE ratings SET rating = ?, photoId = ? WHERE rating_id = ?;", fr.rating, fr.Photo_id, fr.Rating_id)
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
