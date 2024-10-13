package models

import "time"

type User struct {
	ID			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	Email 		string		`json:"email" gorm:"unique"`
	Avatar		string		`json:"avatar"`
	DelBy		bool		`json:"del_by" gorm:"default:false"`
	CreatedAt 	time.Time 	`json:"created_at"`
    UpdatedAt 	time.Time 	`json:"updated_at"`
}