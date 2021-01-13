package memory

import (
	"api-customers/application/model"
	"context"
)

type customerInMemoryRepository struct {
}

func NewCustomerInMemoryRepository() *customerInMemoryRepository {
	return &customerInMemoryRepository{}
}

func (repo customerInMemoryRepository) FindAll(ctx context.Context) ([]model.Customer, error) {
	/*
		tr:=otel.Tracer("")
		ctx,span:=tr.Start(ctx,"")*/

	return []model.Customer{
		{
			Document: "47506306",
			Name:     "Alexander",
			Role:     "Software Architect",
			Salary:   12000,
		},
		{
			Document: "34806205",
			Name:     "Esteban",
			Role:     "Software Developer",
			Salary:   5000,
		},
	}, nil
}
