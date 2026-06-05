package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Age       int       `json:"age"`
	Status    string    `json:"status"`
	AddedAt   time.Time `json:"addedat"`
	AddedBy   string    `json:"addedby"`
	UpdatedAt time.Time `json:"updatedat"`
	UpdatedBy string    `json:"updatedby"`
}