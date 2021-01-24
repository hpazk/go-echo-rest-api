package product

import "github.com/hpazk/go-echo-rest-api/app/models"

type Category struct {
	models.Base
	Name string
}
