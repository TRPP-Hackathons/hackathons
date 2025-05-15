package hackathons

import (
	"context"
	"fmt"

	"hackathons/internal/models"
	"hackathons/internal/ports/repository"
	"hackathons/internal/services"
)

type usersService struct {
	usersRepository repository.UsersRepository
}

func NewUsersService(
	usersRepository repository.UsersRepository,
) services.Users {
	return &usersService{
		usersRepository: usersRepository,
	}
}

func (us *usersService) GetUser(ctx context.Context) (models.User, error) {
	user, err := us.usersRepository.GetUserByID(ctx, services.MockUserID)
	if err != nil {
		return models.User{}, fmt.Errorf("repo get user by id: %w", err)
	}

	return user, nil
}

func (us *usersService) GetParticipants(ctx context.Context) ([]models.Participant, error) {
	participants, err := us.usersRepository.GetParticipants(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get participants: %w", err)
	}

	return participants, nil
}
