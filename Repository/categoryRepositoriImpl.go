package repository

import (
	"context"
	"database/sql"
	"restfull-api/Model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("Implement me")
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("Implement me")
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	panic("Implement me")
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("Implement me")
}
func (repository *CategoryRepositoryImpl) findAll(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category {
	panic("Implement me")
}
