package repository

import (
	"P1_API/helper"
	"P1_API/model"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookID int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from book where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, bookID)
	helper.PanicIfError(errExec)
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name from book"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookID int) (model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name from from book where id=$1"
	resut, errQuery := tx.QueryContext(ctx, SQL, bookID)
	helper.PanicIfError(errQuery)
	defer resut.Close()

	book := model.Book{}

	if resut.Next() {
		err := resut.Scan(&bookID, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("Book id is not found")
	}
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into book(name) value($1)"
	_, err = tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicIfError(err)
}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update book set name=$1 where id=$2"
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(err)
}