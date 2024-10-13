package models

import "database/sql"

type User struct {
    ID        uint         `json:"id" gorm:"primaryKey"`
    Name      string       `json:"name"`
    Email     string       `json:"email" gorm:"unique"`
    Password  string       `json:"password"`
    Avatar    string       `json:"avatar"`
    DelBy     bool         `json:"del_by" gorm:"default:false"`
    CreatedAt sql.NullTime `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
