package services

import (
	"api/cmd/internal/postgresrepo"
	"time"
)

var defaultDBTimeout = 30 * time.Second

type Service struct {
	queries *postgresrepo.Queries
}

func NewService(queries *postgresrepo.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) TimeNow() time.Time {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		// Handle error
		return time.Now().UTC() // fallback to UTC in case of error
	}
	return time.Now().In(loc)
}
