package services

import (
	"api-customers/application/model"
	"api-customers/application/storage"
	"context"
)

type CustomerFinder struct {
	repo storage.CustomerRepository
}

func NewCustomerFinder(repo storage.CustomerRepository) *CustomerFinder {
	return &CustomerFinder{repo: repo}
}

func (service CustomerFinder) GetAll(ctx context.Context) ([]model.Customer, error) {

	return service.repo.FindAll(ctx)
}
