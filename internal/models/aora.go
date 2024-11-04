package models

type Aora struct {
	ID			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	VideoURL	string		`json:"video_url"`
	VideoThumb	string		`json:"video_thumb"`
	Author		string		`json:"author"`
	AuthorPhoto string		`json:"author_photo"`
	Description	string		`json:"description"`
}