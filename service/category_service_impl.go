package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/exception"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/domain"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/web"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository // Inject CategoryRepository for Implement Service
	DB                 *sql.DB
	Validate           *validator.Validate
}

// HTTP Router Impl
func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// Start Validtaion
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	// Start Transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx) // Commit Or Rollback
	// RequestBody
	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	// My Way
	// var categoryResponse domain.Category

	// return web.CategoryResponse(categoryResponse)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// Start Validtaion
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)
	// Start Transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx) // Commit Or Rollback
	// Request Body
	// Find Id
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	// Err For Not Found Error Web Response
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	// Set category Name to request name if id Found
	category.Name = request.Name
	//
	category = service.CategoryRepository.Update(ctx, tx, category)
	// Convert Category To Category Response
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// Start Transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// RequestBody
	// Find Id
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// Err For Not Found Error Web Response
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// Err For Not Found Error Web Response
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	// var categoryResponses []web.CategoryResponse
	// for _, category := range categories {
	// 	categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))

	// }
	return helper.ToCategorResponses(categories)

}
