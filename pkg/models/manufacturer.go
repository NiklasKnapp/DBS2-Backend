package models

import (
	"database/sql"
	"fmt"

	"de.stuttgart.hft/DBS2-Backend/pkg/config"
)

var db *sql.DB

type Manufacturer struct {
	M_id int    `json:"mId"`
	Name string `json:"name" binding:"required"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (m *Manufacturer) CreateManufacturer() (*Manufacturer, error) {
	res, err := db.Exec("INSERT INTO manufacturer (name) VALUES(?);", m.Name)
	if err != nil {
		return nil, fmt.Errorf("CreateManufacturer: %v", err)
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateManufacturer: %v", err)
	}
	m.M_id = int(id)
	return m, nil
}

func GetManufacturer() ([]Manufacturer, error) {
	//Empty struct == [], uninitialized struct returns nil
	//https://stackoverflow.com/questions/56200925/return-an-empty-array-instead-of-null-with-golang-for-json-return-with-gin
	var manufacturers = []Manufacturer{}
	rows, err := db.Query("SELECT m_id, name FROM manufacturer") //manufacturer
	if err != nil {
		return nil, fmt.Errorf("GetManufacturer: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var mans Manufacturer
		if err := rows.Scan(&mans.M_id, &mans.Name); err != nil {
			return nil, fmt.Errorf("GetManufacturer: %v", err)
		}
		manufacturers = append(manufacturers, mans)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetManufacturer: %v", err)
	}
	return manufacturers, nil
}

func GetManufacturerById(MId int64) (*Manufacturer, error) {
	man := &Manufacturer{}
	if err := db.QueryRow("SELECT m_id, name FROM manufacturer WHERE m_id = ?;", MId).Scan(&man.M_id, &man.Name); err != nil {
		return nil, fmt.Errorf("GetManufacturerById: %v", err)
	}
	return man, nil
}

func (m *Manufacturer) UpdateManufacturer() (*Manufacturer, error) {
	_, err := db.Exec("UPDATE manufacturer SET name = ? WHERE m_id = ?;", m.Name, m.M_id)
	if err != nil {
		return nil, fmt.Errorf("GetManufacturerById: %v", err)
	}
	return m, nil
}

func DeleteManufacturer(Mid int64) (*Manufacturer, error) {
	man, _ := GetManufacturerById(Mid)
	_, err := db.Exec("DELETE FROM manufacturer WHERE m_id = ?;", Mid)
	if err != nil {
		return nil, fmt.Errorf("DeleteManufacturer: %v", err)
	}
	return man, nil
}
