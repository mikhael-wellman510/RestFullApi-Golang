package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"restfull-api/Model/domain"
	"restfull-api/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	SQL := "insert into category (name) values ($1) returning id"

	var id int

	err := tx.QueryRow(SQL, category.Name).Scan(&id)

	helper.PanicIfErr(err)
	category.Id = id
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = $1 where id = $2"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)

	helper.PanicIfErr(err)

	return category

}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id=$1"
	fmt.Println(category)

	res, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)

	// Optional: Check if the query affected any rows (i.e., if the category was deleted)
	rowsAffected, err := res.RowsAffected()
	helper.PanicIfErr(err)

	// If no rows were affected, return an error indicating the category was not found
	fmt.Println(rowsAffected)
	// Return nil if the deletion was successful

}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id , name from category where id = $1"

	// Mencari data menggunakan Query Context
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErr(err)

	// Wajib di close
	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)

	defer rows.Close()
	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		categories = append(categories, category)
	}

	return categories

}
