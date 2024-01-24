package main

import (
	"encoding/json"
	"os"
)

type Bread struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Origin      string   `json:"origin"`
	Ingredients []string `json:"ingredients"`
}

type BreadStore struct {
	breads []Bread
}

func (bs *BreadStore) LoadData(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &bs.breads)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BreadStore) GetAll() []Bread {
	return bs.breads
}
