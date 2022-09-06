package service

import (
	"context"
	"fmt"
	"practice/internal/models"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetFriendsById(id uint) ([]string, error)
	AddFriend(ctx context.Context, f models.Friendship) (name1, name2 string, err error)
	DeleteUser(ctx context.Context, id *uint) error
	GetUserById(ctx context.Context, id *uint) (*models.User, error)
	CreateUser(ctx context.Context, u *models.User) (uint, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type service struct {
	log        *logrus.Logger
	repository Repository
}

//fabric funcion
func NewService(logger *logrus.Logger, r Repository) *service {
	return &service{
		log:        logger,
		repository: r,
	}
}

//create user
func (s *service) CreateUser(ctx context.Context, u *models.User) (uint, error) {
	s.log.Info("create user in service")
	return s.repository.CreateUser(ctx, u)
}

//get all users
func (s *service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repository.GetAllUsers(ctx)
}

//get by id
func (s *service) GetUserById(ctx context.Context, id *uint) (u *models.User, err error) {
	return s.repository.GetUserById(ctx, id)
}

//delete user by id
func (s *service) DeleteUser(ctx context.Context, id *uint) error {
	fmt.Println("id:", id)
	return s.repository.DeleteUser(ctx, id)
}
func (s *service) AddFriend(ctx context.Context, f models.Friendship) (name1, name2 string, err error) {
	return s.repository.AddFriend(ctx, f)
}
func (s *service) GetFriendsById(id uint) ([]string, error) {
	return s.repository.GetFriendsById(id)
}
