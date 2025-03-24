package services

import (
	"github.com/Lets-Golang/Travel-Booking/user-service/models"
	"github.com/Lets-Golang/Travel-Booking/user-service/repositories"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Create(dto models.UserDTO) (models.UserDTO, error) {
    user := models.UserEntity{Name: dto.Name, Email: dto.Email}
    err := s.repo.Create(&user)
    return models.UserDTO{ID: user.ID, Name: user.Name, Email: user.Email}, err
}

func (s *UserService) GetByID(id int) (models.UserDTO, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return models.UserDTO{}, err
    }
    return models.UserDTO{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}
