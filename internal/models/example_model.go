package models

import "gorm.io/gorm"

type Example struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e *Example) Create(db *gorm.DB) error {
	return db.Create(e).Error
}

func (e *Example) Update(db *gorm.DB) error {
	return db.Save(e).Error
}

func (e *Example) Delete(db *gorm.DB) error {
	return db.Delete(e).Error
}

func GetExampleByID(db *gorm.DB, id int) (*Example, error) {
	var example Example
	err := db.First(&example, id).Error
	return &example, err
}

func GetAllExamples(db *gorm.DB) ([]Example, error) {
	var examples []Example
	err := db.Find(&examples).Error
	return examples, err
}
