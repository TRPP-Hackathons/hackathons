package hackathons

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"hackathons/internal/infrastructure/database/postgres"
	"hackathons/internal/models"
	"hackathons/internal/ports/repository"
)

const (
	usersTable                 = "users"
	participantsTable          = "participants"
	participantsTableWithAlias = participantsTable + " p"
)

type usersRepository struct {
	db *postgres.Postgres
}

func NewUsersRepository(
	db *postgres.Postgres,
) repository.UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) GetUserByID(ctx context.Context, userID int) (models.User, error) {
	qb := r.db.Builder.Select(
		"id",
		"first_name",
		"last_name",
		"image_url",
	).
		From(usersTable).
		Where(squirrel.Eq{"id": userID})

	query, args, err := qb.ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("to sql %w", err)
	}

	var user models.User
	if err = r.db.SqlxDB().GetContext(ctx, &user, query, args...); err != nil {
		return models.User{}, fmt.Errorf("select %w", err)
	}

	return user, nil
}

func (r *usersRepository) GetParticipants(ctx context.Context) ([]models.Participant, error) {
	qb := r.db.Builder.Select(
		"u.id",
		"u.first_name",
		"u.last_name",
		"u.image_url",
		"p.role",
	).
		From(participantsTableWithAlias).
		Join("users u ON p.user_id = u.id")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	var participants []models.Participant
	if err = r.db.SqlxDB().SelectContext(ctx, &participants, query, args...); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	return participants, nil
}
