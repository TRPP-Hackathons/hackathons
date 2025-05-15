package hackathons

import (
	"context"
	"fmt"

	"hackathons/internal/infrastructure/database/postgres"
	"hackathons/internal/models"
	"hackathons/internal/ports/repository"
)

const (
	hackathonsTable = "hackathons"
)

type hackathonsRepository struct {
	db *postgres.Postgres
}

func NewHackathonsRepository(
	db *postgres.Postgres,
) repository.HackathonsRepository {
	return &hackathonsRepository{
		db: db,
	}
}

func (r *hackathonsRepository) GetHackathons(ctx context.Context) ([]models.Hackathon, error) {
	qb := r.db.Builder.Select(
		"id",
		"name",
		"money",
		"participant_min",
		"participant_max",
		"image_url",
	).
		From(hackathonsTable)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql %w", err)
	}

	var hackathons []models.Hackathon
	if err = r.db.SqlxDB().SelectContext(ctx, &hackathons, query, args...); err != nil {
		return nil, fmt.Errorf("select %w", err)
	}

	return hackathons, nil
}
