package models

import "fmt"

// Roll type
type RollType struct {
	Type_id   int    `json:"typeId"`
	StockName string `json:"stockName" binding:"required"`
	Format    string `json:"format" binding:"required"`
	M_id      int    `json:"mId" binding:"required"`
}

// Create roll type in DB
func (rt *RollType) CreateRollType() (*RollType, error) {
	
	// Check for existing manufacturer
	m_id, _ := GetManufacturerById(int64(rt.M_id))
	if m_id == nil {
		return nil, fmt.Errorf("CreateRollType: Manufacturer with m_id %v does not exist", rt.M_id)
	}

	// Run query
	res, err := db.Exec("INSERT INTO roll_type (stock_name, size, m_id) VALUES(?, ?, ?);", rt.StockName, rt.Format, rt.M_id)
	if err != nil {
		return nil, fmt.Errorf("CreateRollType: %v", err)
	}

	// Get ID
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateRollType: %v", err)
	}
	rt.Type_id = int(id)
	return rt, nil
}

// Get roll types from DB
func GetRollType() ([]RollType, error) {

	// List of roll types
	var rollTypes = []RollType{}

	// Run query
	rows, err := db.Query("SELECT type_id, stock_name, size, m_id FROM roll_type;")
	if err != nil {
		return nil, fmt.Errorf("GetRollType: %v", err)
	}
	defer rows.Close()

	// Extract values
	for rows.Next() {
		var types RollType
		if err := rows.Scan(&types.Type_id, &types.StockName, &types.Format, &types.M_id); err != nil {
			return nil, fmt.Errorf("GetRollType: %v", err)
		}
		rollTypes = append(rollTypes, types)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetRollType: %v", err)
	}
	return rollTypes, nil
}

// Get roll type by ID from DB
func GetRollTypeById(tId int64) (*RollType, error) {
	roll := &RollType{}
	if err := db.QueryRow("SELECT type_id, stock_name, size, m_id FROM roll_type WHERE type_id = ?;", tId).Scan(&roll.Type_id, &roll.StockName, &roll.Format, &roll.M_id); err != nil {
		return nil, fmt.Errorf("GetRollTypeById: %v", err)
	}
	return roll, nil
}

// Update roll type in DB
func (rt *RollType) UpdateRollType() (*RollType, error) {
	_, err := db.Exec("UPDATE roll_type SET stock_name = ?, size = ?, m_id = ? WHERE type_id = ?;", rt.StockName, rt.Format, rt.M_id, rt.Type_id)
	if err != nil {
		return nil, fmt.Errorf("UpdateRollType: %v", err)
	}
	return rt, nil
}

// Delete roll type from DB
func DeleteRollType(tId int64) (*RollType, error) {
	roll, _ := GetRollTypeById(tId)
	_, err := db.Exec("DELETE FROM roll_type WHERE type_id = ?;", tId)
	if err != nil {
		return nil, fmt.Errorf("DeleteRollType: %v", err)
	}
	return roll, nil
}
