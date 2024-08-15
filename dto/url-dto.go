package dto

import "github.com/go-playground/validator/v10"

type UrlDTO struct {
	Url string `json:"url" validate:"required,url,min=5"`
}

// validate dto with the tags validation rules
func (dto *UrlDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
