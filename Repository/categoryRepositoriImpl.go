package repository

import (
	"context"
	"database/sql"
	"errors"
	"restfull-api/Model/domain"
	"restfull-api/helper"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)

	helper.PanicIfErr(err)

	id, err := result.LastInsertId()

	helper.PanicIfErr(err)

	category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)

	helper.PanicIfErr(err)

	return category

}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)

}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id,name, from category where id = ?"

	// Mencari data menggunakan Query Context
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErr(err)

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)

		return category, nil
	} else {
		return category, errors.New("Category is not Found")
	}

}
func (repository *CategoryRepositoryImpl) findAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		categories = append(categories, category)
	}

	return categories

}
