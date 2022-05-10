package models

import "fmt"

type RollType struct {
	Type_id   int    `json:"typeId"`
	StockName string `json:"stockName" binding:"required"`
	Format    string `json:"format" binding:"required"`
	M_id      int    `json:"mId" binding:"required"`
}

func (rt *RollType) CreateRollType() (*RollType, error) {
	//Validate if Manufacturer with M_id exists
	m_id, _ := GetManufacturerById(int64(rt.M_id))
	if m_id == nil {
		return nil, fmt.Errorf("CreateRollType: Manufacturer with m_id %v does not exist", rt.M_id)
	}
	res, err := db.Exec("INSERT INTO roll_type (stock_name, size, m_id) VALUES(?, ?, ?);", rt.StockName, rt.Format, rt.M_id)
	if err != nil {
		return nil, fmt.Errorf("CreateRollType: %v", err)
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, fmt.Errorf("CreateRollType: %v", err)
	}
	rt.Type_id = int(id)
	return rt, nil
}

func GetRollType() ([]RollType, error) {
	var rollTypes = []RollType{}
	rows, err := db.Query("SELECT type_id, stock_name, size, m_id FROM roll_type;")
	if err != nil {
		return nil, fmt.Errorf("GetRollType: %v", err)
	}
	defer rows.Close()
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

func GetRollTypeById(tId int64) (*RollType, error) {
	roll := &RollType{}
	if err := db.QueryRow("SELECT type_id, stock_name, size, m_id FROM roll_type WHERE type_id = ?;", tId).Scan(&roll.Type_id, &roll.StockName, &roll.Format, &roll.M_id); err != nil {
		return nil, fmt.Errorf("GetRollTypeById: %v", err)
	}
	return roll, nil
}

func (rt *RollType) UpdateRollType() (*RollType, error) {
	_, err := db.Exec("UPDATE roll_type SET stock_name = ?, size = ?, m_id = ? WHERE type_id = ?;", rt.StockName, rt.Format, rt.M_id, rt.Type_id)
	if err != nil {
		return nil, fmt.Errorf("UpdateRollType: %v", err)
	}
	return rt, nil
}

func DeleteRollType(tId int64) (*RollType, error) {
	roll, _ := GetRollTypeById(tId)
	_, err := db.Exec("DELETE FROM roll_type WHERE type_id = ?;", tId)
	if err != nil {
		return nil, fmt.Errorf("DeleteRollType: %v", err)
	}
	return roll, nil
}
