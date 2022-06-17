package models

import (
	"fmt"
)

// Manufacturer
type Manufacturer struct {
	M_id int    `json:"mId"`
	Name string `json:"name" binding:"required"`
}

// Create manufacturer in DB
func (m *Manufacturer) CreateManufacturer() (*Manufacturer, error) {

	// Run query
	res, err := db.Exec("INSERT INTO manufacturer (name) VALUES(?);", m.Name)
	if err != nil {
		return nil, fmt.Errorf("CreateManufacturer: %v", err)
	}

	// Get ID
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateManufacturer: %v", err)
	}
	m.M_id = int(id)
	return m, nil
}

Get manufacturers from DB
func GetManufacturer() ([]Manufacturer, error) {

	// List of manufacturers
	var manufacturers = []Manufacturer{}

	// Run query
	rows, err := db.Query("SELECT m_id, name FROM manufacturer") //manufacturer
	if err != nil {
		return nil, fmt.Errorf("GetManufacturer: %v", err)
	}
	defer rows.Close()

	// Extract values
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

// Get manufacturer by ID from DB
func GetManufacturerById(MId int64) (*Manufacturer, error) {
	man := &Manufacturer{}
	if err := db.QueryRow("SELECT m_id, name FROM manufacturer WHERE m_id = ?;", MId).Scan(&man.M_id, &man.Name); err != nil {
		return nil, fmt.Errorf("GetManufacturerById: %v", err)
	}
	return man, nil
}

// Update manufacturers in DB
func (m *Manufacturer) UpdateManufacturer() (*Manufacturer, error) {
	_, err := db.Exec("UPDATE manufacturer SET name = ? WHERE m_id = ?;", m.Name, m.M_id)
	if err != nil {
		return nil, fmt.Errorf("UpdateManufacturer: %v", err)
	}
	return m, nil
}

// Delete manufacturers from DB
func DeleteManufacturer(Mid int64) (*Manufacturer, error) {
	man, _ := GetManufacturerById(Mid)
	_, err := db.Exec("DELETE FROM manufacturer WHERE m_id = ?;", Mid)
	if err != nil {
		return nil, fmt.Errorf("DeleteManufacturer: %v", err)
	}
	return man, nil
}
