package service

import (
	"aditditkudit/learn-go-restful-api/helper"
	"aditditkudit/learn-go-restful-api/model/domain"
	"aditditkudit/learn-go-restful-api/model/web"
	"aditditkudit/learn-go-restful-api/repository"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	category.Name = request.Name

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfErr(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfErr(err)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
