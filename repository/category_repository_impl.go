package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/fahmikudo/example-restful-api/helper"
	"github.com/fahmikudo/example-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepositoryImpl(DB *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: DB}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, category domain.Category) domain.Category {
	SQL := "INSERT INTO category (name) VALUES (?)"

	tx, err := c.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	category.Id = int(id)
	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"

	tx, err := c.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	_, err = tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfErr(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, category domain.Category) {
	SQL := "DELETE FROM category WHERE id = ?"

	tx, err := c.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	_, err = tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"

	tx, err := c.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErr(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfErr(err)
	}(rows)

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		return category, nil
	} else {
		return category, errors.New("category not found.")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context) []domain.Category {
	SQL := "SELECT id, name FROM category"

	tx, err := c.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfErr(err)
	}(rows)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		categories = append(categories, category)
	}
	return categories
}
