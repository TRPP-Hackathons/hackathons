package services

import (
	"context"

	"hackathons/internal/models"
)

type Users interface {
	GetUser(ctx context.Context) (models.User, error)
	GetParticipants(ctx context.Context) ([]models.Participant, error)
}

type Hackathons interface {
	GetHackathons(ctx context.Context) ([]models.Hackathon, error)
}
