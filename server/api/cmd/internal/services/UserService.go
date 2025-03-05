package services

import (
	"api/cmd/internal/postgresrepo"
	"context"
)

func (s *Service) CreateUser(userParams *postgresrepo.CreateUserParams) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	err := s.queries.CreateUser(ctx, *userParams)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserByEmail(email string) (*postgresrepo.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) GetUserByID(id int32) (*postgresrepo.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) CheckIfUserExistsByEmail(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	count, err := s.queries.CountUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Service) CheckIfUserExistsByUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	count, err := s.queries.CountUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Service) UpdateUserPasswordByID(params *postgresrepo.UpdateUserPasswordByIDParams) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	err := s.queries.UpdateUserPasswordByID(ctx, *params)
	if err != nil {
		return err
	}

	return nil
}
