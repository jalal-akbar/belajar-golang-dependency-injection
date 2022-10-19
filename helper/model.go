package helper

import (
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/domain"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	// Mas Eko Way
	// return web.CategoryResponse{
	// 	Id:   category.Id,
	// 	Name: category.Name,
	// }
	// My Way
	categoryResponse := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}

func ToCategorResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
