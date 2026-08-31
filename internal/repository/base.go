package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("data not found")
)

// BaseRepository provides common CRUD operations for all repositories
type BaseRepository[T any] struct {
	Db *gorm.DB
}

// DB returns the underlying database instance
func (r *BaseRepository[T]) DB() *gorm.DB {
	return r.Db
}

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	Get(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]T, error)
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		Db: db,
	}
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.Db.Create(entity).Error
}

func (r *BaseRepository[T]) Get(ctx context.Context, id string) (*T, error) {
	var entity T
	err := r.Db.First(&entity, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.Db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id string) error {
	return r.Db.Delete(new(T), "id=?", id).Error
}

func (r *BaseRepository[T]) List(ctx context.Context, offset, limit int) ([]T, error) {
	var entities []T
	if err := r.Db.WithContext(ctx).Offset(offset).Limit(limit).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}
