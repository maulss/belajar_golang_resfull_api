package helper

import (
	"belajar_golang_resrful_api/model/domain"
	"belajar_golang_resrful_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
