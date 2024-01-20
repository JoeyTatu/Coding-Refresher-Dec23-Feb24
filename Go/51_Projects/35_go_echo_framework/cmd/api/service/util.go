package service

import (
	"encoding/json"
	"os"
)

type Data struct {
	UserID int    `json:"user_id"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Payload struct {
	Data []Data `json:"data"`
}

func raw() ([]Data, error) {
	r, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var payload Payload
	err = json.Unmarshal(r, &payload.Data)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetByID(id int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}

	if id > len(data) {
		res := make([]string, 0)
		return res, nil
	}

	return data[id], nil
}
