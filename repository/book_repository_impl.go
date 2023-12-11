package repository

import (
	"P1_API/helper"
	"P1_API/model"
	"context"
	"database/sql"
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
	_, errExec := tx.ExecContext(ctx,SQL,bookID)
	helper.PanicIfError(errExec)
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	panic("unimplemented")
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookID int) (model.Book, error) {
	panic("unimplemented")
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	panic("unimplemented")
}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	panic("unimplemented")
}


