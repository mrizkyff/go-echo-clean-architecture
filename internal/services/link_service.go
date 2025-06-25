package services

import (
	"go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/repositories"
	"go-echo-clean-architecture/internal/utils"

	"github.com/google/uuid"
)

type LinkService interface {
	GetLinkById(id uuid.UUID, userID uuid.UUID) (*models.Link, error)
	GetAllLinkWithPagination(pagination utils.PaginationType) ([]*models.Link, int64, error)
	CreateLink(link *models.Link) (*models.Link, error)
	UpdateLink(id uuid.UUID, link *models.Link) (*models.Link, error)
	DeleteLink(id uuid.UUID) (*models.Link, error)
}

type LinkServiceImpl struct {
	linkRepository repositories.LinkRepository
	userRepository repositories.UserRepository
}

func NewLinkServiceImpl(linkRepository repositories.LinkRepository, userRepository repositories.UserRepository) *LinkServiceImpl {
	return &LinkServiceImpl{linkRepository: linkRepository, userRepository: userRepository}
}

func (l *LinkServiceImpl) GetLinkById(id uuid.UUID, userID uuid.UUID) (*models.Link, error) {
	findById, err := l.linkRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if findById.UserID != userID {
		return nil, errors.NewNotFoundError("Link not found in this account")
	}

	return findById, nil
}

func (l *LinkServiceImpl) GetAllLinkWithPagination(pagination utils.PaginationType) ([]*models.Link, int64, error) {
	allWithPagination, totalRecords, err := l.linkRepository.FindAllWithPagination(pagination)
	if err != nil {
		return nil, 0, err
	}

	return allWithPagination, totalRecords, nil
}

func (l *LinkServiceImpl) CreateLink(link *models.Link) (*models.Link, error) {
	// Get the user
	_, err := l.userRepository.FindById(link.CreatedBy)
	if err != nil {
		return nil, err
	}

	// Assign user to link
	link.UserID = link.CreatedBy

	create, err := l.linkRepository.Create(link)
	if err != nil {
		return nil, err
	}

	linkById, err := l.GetLinkById(create.ID, create.UserID)
	if err != nil {
		return nil, err
	}

	return linkById, nil
}

func (l *LinkServiceImpl) UpdateLink(id uuid.UUID, link *models.Link) (*models.Link, error) {
	_, err := l.linkRepository.FindById(id)
	if err != nil {
		return nil, errors.NewNotFoundError("Link not found")
	}

	update, err := l.linkRepository.Update(id, link)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to update link")
	}

	linkById, err := l.GetLinkById(update.ID, update.UserID)
	if err != nil {
		return nil, err
	}

	return linkById, err
}

func (l *LinkServiceImpl) DeleteLink(id uuid.UUID) (*models.Link, error) {
	found, err := l.linkRepository.FindById(id)
	if err != nil {
		return nil, errors.NewNotFoundError("Link not found")
	}

	_, err = l.GetLinkById(found.ID, found.UserID)
	if err != nil {
		return nil, err
	}

	err = l.linkRepository.Delete(id)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to delete link")
	}

	return found, nil
}
