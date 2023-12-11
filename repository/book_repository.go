package repository

import (
	"P1_API/model"
	"context"
)

type BookRepository interface {
	Save(ctx context.Context, book model.Book)
	Update(ctx context.Context, book model.Book)
	Delete(ctx context.Context, bookID int)
	FindById(ctx context.Context, bookID int)(model.Book, error)
	FindAll(ctx context.Context) []model.Book
}