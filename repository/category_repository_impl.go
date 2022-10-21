package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/domain"
)

type CategoryRepositoryImpl struct {
}

// HTTP Router Impl
func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repo *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name)VALUES(?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	res, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(res)

	return category
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "UPDATE category SET name=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)

}
func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}

}
func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories = []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
