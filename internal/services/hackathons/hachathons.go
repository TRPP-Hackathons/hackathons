package hackathons

import (
	"context"
	"fmt"

	"hackathons/internal/models"
	"hackathons/internal/ports/repository"
	"hackathons/internal/services"
)

type hackathonsService struct {
	hackathonsRepository repository.HackathonsRepository
}

func NewHackathonsService(
	hackathonsRepository repository.HackathonsRepository,
) services.Hackathons {
	return &hackathonsService{
		hackathonsRepository: hackathonsRepository,
	}
}

func (us *hackathonsService) GetHackathons(ctx context.Context) ([]models.Hackathon, error) {
	hackathons, err := us.hackathonsRepository.GetHackathons(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get hachathons: %w", err)
	}

	return hackathons, nil
}
