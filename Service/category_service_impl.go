package service

import (
	"context"
	"database/sql"
	"restfull-api/Model/domain"
	"restfull-api/Model/web"
	repository "restfull-api/Repository"
	"restfull-api/helper"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	// Injeksi Repository
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	// Validator package
	Validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfErr(errValidate)
	tx, err := service.DB.Begin()

	helper.PanicIfErr(err)
	// Kalau ada error transaksi . dia akan rollback
	defer helper.CommitOrRollback(tx)

	// Tidak pakai ID . karena autogenerate
	category := domain.Category{
		Name: request.Name,
	}

	// Panggil Repositori nya
	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfErr(errValidate)
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)

	// Find Id , ada atau tidak
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	// Kalau ada . assign (Updated)
	category.Name = request.Name

	// Panggil Repositori
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)

	// cari dulu id yang mau di hapus
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	helper.PanicIfErr(err)
	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)

	// Find id
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	helper.PanicIfErr(err)

	// Transalate categoryResponse
	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)

	// Category tidak menggunakan err , karena return nya hanya array
	category := service.CategoryRepository.FindAll(ctx, tx)

	data := []web.CategoryResponse{}

	for _, val := range category {
		// ubah dulu ke categoryResponse
		res := web.CategoryResponse{}
		res.Id = val.Id
		res.Name = val.Name
		data = append(data, res)
	}

	return data

}
