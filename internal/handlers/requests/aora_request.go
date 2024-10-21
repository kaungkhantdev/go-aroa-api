package requests

type AoraCreateRequest struct {
	Name		string	`json:"name" validate:"required,min=3"`
	VideoURL	string	`json:"video_url" validate:"required"`
	VideoThumb	string	`json:"video_thumb" validate:"required"`
	Author		string	`json:"author" validate:"required"`
	AuthorPhoto string	`json:"author_photo" validate:"required"`
	Description	string	`json:"description" validate:"required"`
}