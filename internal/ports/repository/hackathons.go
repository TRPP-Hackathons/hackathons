package repository

import (
	"context"

	"hackathons/internal/models"
)

type UsersRepository interface {
	GetUserByID(ctx context.Context, userID int) (models.User, error)
	GetParticipants(ctx context.Context) ([]models.Participant, error)
}

type HackathonsRepository interface {
	GetHackathons(ctx context.Context) ([]models.Hackathon, error)
}
