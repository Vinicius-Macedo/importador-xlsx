package services

import (
	"api/cmd/internal/postgresrepo"
	"context"
)

func (s *Service) GetAllCustomers() ([]postgresrepo.GetAllCustomersRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	rows, err := s.queries.GetAllCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Service) GetAllResources() ([]postgresrepo.GetAllResourcesRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	rows, err := s.queries.GetAllResources(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Service) GetAllCategories() ([]postgresrepo.GetCategoriesRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	rows, err := s.queries.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
