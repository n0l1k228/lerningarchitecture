package service

import (
	"context"
	"fmt"
	"myavito/internal/domain"
)

type UserRepository interface {
	SaveUser(ctx context.Context, name, phone_number string) (int, error)
}

type ProductRepository interface {
	SaveProduct(ctx context.Context, user_id, price int, title, description, category string) error
	SearchProduct(ctx context.Context, title string) ([]domain.Product, error)
}

type UserService struct {
	repo UserRepository
}

type ProductService struct {
	repo ProductRepository
}

func NewServiceUser(us UserRepository) *UserService {
	return &UserService{
		repo: us,
	}
}

func NewProductService(us ProductRepository) *ProductService {
	return &ProductService{
		repo: us,
	}
}

func (us *UserService) ServiceCreateUser(ctx context.Context, name, phone_number string) (int, error) {
	id, err := us.repo.SaveUser(ctx, name, phone_number)
	if err != nil {
		return -1, fmt.Errorf("failed save user in db: %w", err)
	}
	return id, nil
}

func (us *ProductService) ServiceCreateProduct(ctx context.Context, user_id, price int, title, description, category string) error {
	err := us.repo.SaveProduct(ctx, user_id, price, title, description, category)
	if err != nil {
		return fmt.Errorf("failed save user in db: %w", err)
	}
	return nil
}

func (us *ProductService) ServiceSearchProduct(ctx context.Context, title string) ([]domain.Product, error) {
	prod, err := us.repo.SearchProduct(ctx, title)
	if err != nil {
		return nil, fmt.Errorf("failed save user in db: %w", err)
	}
	return prod, nil
}
