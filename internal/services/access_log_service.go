package services

import (
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/repositories"
	"go-echo-clean-architecture/internal/utils"
)

type AccessLogService struct {
	accessLogRepositories repositories.AccessLogRepository
}

func NewAccessLogService(accessLogRepositories repositories.AccessLogRepository) *AccessLogService {
	return &AccessLogService{accessLogRepositories: accessLogRepositories}
}

func (a *AccessLogService) Create(log *models.AccessLog) (*models.AccessLog, error) {
	createdLog, err := a.accessLogRepositories.Create(log)
	if err != nil {
		return nil, err
	}

	return createdLog, nil
}

func (a *AccessLogService) GetAll(pagination utils.PaginationType) ([]*models.AccessLog, int64, error) {
	result, i, err := a.accessLogRepositories.FindAllWithPagination(pagination)
	if err != nil {
		return nil, 0, err
	}

	return result, i, err
}
