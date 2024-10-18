package helper

import (
	"restfull-api/Model/domain"
	"restfull-api/Model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {

	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

}
