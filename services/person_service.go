package services

import (
	"context"
	"marrywith-gin-api/models"
)

type PersonService interface {
	CreatePerson(ctx context.Context, person *models.Person) error
	GetPersons(ctx context.Context) ([]models.Person, error)
}

type personService struct {
	personRepo models.PersonRepository
}

func NewPersonService(repo models.PersonRepository) PersonService {
	return &personService{personRepo: repo}
}

func (s *personService) CreatePerson(ctx context.Context, person *models.Person) error {
	return s.personRepo.Create(ctx, person)
}

func (s *personService) GetPersons(ctx context.Context) ([]models.Person, error) {
	return s.personRepo.GetAll(ctx)
}
