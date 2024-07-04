package service

import (
	"context"
	"github.com/fahmikudo/example-restful-api/exception"
	"github.com/fahmikudo/example-restful-api/helper"
	"github.com/fahmikudo/example-restful-api/model/domain"
	"github.com/fahmikudo/example-restful-api/model/web"
	"github.com/fahmikudo/example-restful-api/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, Validate: validate}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {

	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, category)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {

	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	categories := service.CategoryRepository.FindAll(ctx)

	return helper.ToCategoryResponses(categories)
}
