package storage

import (
	"api-customers/application/model"
	"context"
)

type  CustomerRepository interface {
	FindAll(ctx context.Context) ([]model.Customer, error)
}
