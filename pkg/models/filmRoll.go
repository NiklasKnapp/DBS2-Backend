package models

import (
	"fmt"
)

type FilmRoll struct {
	Roll_id     int     `json:"rollId"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Type_id     int     `json:"typeId" binding:"required"`
	Rating      float32 `json:"rating"`
}

func (fr *FilmRoll) CreateRollType() (*FilmRoll, error) {
	type_id, _ := GetRollTypeById(int64(fr.Type_id))
	if type_id == nil {
		return nil, fmt.Errorf("CreateRollType: RollType with type_id %v does not exist", fr.Type_id)
	}
	res, err := db.Exec("INSERT INTO film_rolls (title, text, type_id) VALUES(?, ?, ?);", fr.Title, fr.Description, fr.Type_id)
	if err != nil {
		return nil, fmt.Errorf("CreateRollType: %v", err)
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateRollType: %v", err)
	}
	fr.Roll_id = int(id)
	return fr, nil
}

func GetFilmRoll() ([]FilmRoll, error) {
	var filmRolls = []FilmRoll{}
	rows, err := db.Query("SELECT roll_id, title, text, type_id, rating FROM film_rolls;")
	if err != nil {
		return nil, fmt.Errorf("GetFilmRoll: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var rolls FilmRoll
		if err := rows.Scan(&rolls.Roll_id, &rolls.Title, &rolls.Description, &rolls.Type_id, &rolls.Rating); err != nil {
			return nil, fmt.Errorf("GetFilmRoll: %v", err)
		}
		filmRolls = append(filmRolls, rolls)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFilmRoll: %v", err)
	}
	return filmRolls, nil
}

func GetFilmRollById(rId int64) (*FilmRoll, error) {
	roll := &FilmRoll{}
	if err := db.QueryRow("SELECT roll_id, title, text, type_id, rating FROM film_rolls WHERE roll_id = ?;", rId).Scan(&roll.Roll_id, &roll.Title, &roll.Description, &roll.Type_id, &roll.Rating); err != nil {
		return nil, fmt.Errorf("GetFilmRollById: %v", err)
	}
	return roll, nil
}

func (fr *FilmRoll) UpdateFilmRoll() (*FilmRoll, error) {
	_, err := db.Exec("UPDATE film_rolls SET title = ?, text = ?, type_id = ?, rating = ? WHERE roll_id = ?;", fr.Title, fr.Description, fr.Type_id, fr.Roll_id, fr.Rating)
	if err != nil {
		return nil, fmt.Errorf("UpdateFilmRoll: %v", err)
	}
	return fr, nil
}

func DeleteFilmRoll(rId int64) (*FilmRoll, error) {
	roll, _ := GetFilmRollById(rId)
	_, err := db.Exec("DELETE FROM film_rolls WHERE roll_id = ?;", rId)
	if err != nil {
		return nil, fmt.Errorf("DeleteFilmRoll: %v", err)
	}
	return roll, nil
}
