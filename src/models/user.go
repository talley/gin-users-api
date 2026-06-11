package models

import "time"

type User struct {
	ID     uint   `json:"id" gorm:"column:id;primaryKey"`
	Name   string `json:"name" gorm:"column:name"`
	Age    int    `json:"age" gorm:"column:age"`
	Status string `json:"status" gorm:"column:status"`

	AddedAt time.Time `json:"addedat" gorm:"column:addedat"`
	AddedBy string    `json:"addedby" gorm:"column:addedby"`

	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
	UpdatedBy string    `json:"updatedby" gorm:"column:updatedby"`
}

func (User) TableName() string {
	return "users"
}
