package services

import (
	"deteleng-backend/internal/models"
	"deteleng-backend/internal/repository"
	"errors"
)

type ServiceService struct {
	repo repository.Repository
}

func NewServiceService(repo repository.Repository) *ServiceService {
	return &ServiceService{repo: repo}
}

func (s *ServiceService) GetAllServices() ([]models.Service, error) {
	return s.repo.GetAllServices()
}

func (s *ServiceService) GetAvailableServices() ([]models.Service, error) {
	allServices, err := s.repo.GetAllServices()
	if err != nil {
		return nil, err
	}

	// Filter only available services
	available := make([]models.Service, 0)
	for _, svc := range allServices {
		if svc.Available {
			available = append(available, svc)
		}
	}
	return available, nil
}

func (s *ServiceService) GetService(id string) (*models.Service, error) {
	return s.repo.GetService(id)
}

func (s *ServiceService) CreateService(service *models.Service) error {
	return s.repo.CreateService(service)
}

func (s *ServiceService) UpdateService(service *models.Service) (*models.Service, error) {
	existing, err := s.repo.GetService(service.ID)
	if err != nil {
		return nil, errors.New("service not found")
	}

	// Update fields
	if service.Name != "" {
		existing.Name = service.Name
	}
	if service.Description != "" {
		existing.Description = service.Description
	}
	if service.Price > 0 {
		existing.Price = service.Price
	}
	if service.Duration > 0 {
		existing.Duration = service.Duration
	}
	existing.Available = service.Available

	if err := s.repo.UpdateService(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *ServiceService) DeleteService(id string) error {
	return s.repo.DeleteService(id)
}
