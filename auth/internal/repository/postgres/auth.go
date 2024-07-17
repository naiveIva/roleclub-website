package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"auth/internal/repository"
	"auth/models"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (auth *AuthRepository) CreateUser(user *models.User) error {
	const fn = "repository.postgres.auth.CreateUser"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`INSERT INTO %s 
			(uuid, first_name, last_name, father_name,
			tel_number, password_hash, is_hse_student)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		usersTable,
	))
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	_, err = stmt.Exec(
		NewUUID(), user.FirstName, user.LastName, user.FatherName,
		user.TelNumber, user.Password, user.IsHSEStudent)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%s: %v", fn, repository.ErrUserAlreadyExists)
		}
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func (auth *AuthRepository) GetUser(telNumber, password string) (*models.User, error) {
	const fn = "repository.postgres.auth.GetUser"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`SELECT * FROM  %s
		WHERE tel_number=$1 AND password_hash=$2`,
		usersTable,
	))
	if err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	usr := &models.User{}
	err = stmt.QueryRow(telNumber, password).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.FatherName,
		&usr.TelNumber,
		&usr.Password,
		&usr.IsHSEStudent,
		&usr.IsBanned,
		&usr.PlayedGames,
		&usr.ConductedGames,
		&usr.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %v", fn, repository.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %v", fn, err)
	}
	return usr, nil
}

func (auth *AuthRepository) GetUserByUUID(uuid uuid.UUID) (*models.User, error) {
	const fn = "repository.postgres.auth.GetUser"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`SELECT * FROM  %s WHERE uuid = $1`,
		usersTable,
	))
	if err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	usr := &models.User{}
	err = stmt.QueryRow(uuid).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.FatherName,
		&usr.TelNumber,
		&usr.Password,
		&usr.IsHSEStudent,
		&usr.IsBanned,
		&usr.PlayedGames,
		&usr.ConductedGames,
		&usr.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %v", fn, repository.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %v", fn, err)
	}
	return usr, nil
}

func (auth *AuthRepository) IncrementPlayedGames(uuid uuid.UUID) error {
	const fn = "repository.postgres.auth.IncrementPlayedGames"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`UPDATE %s SET played_games = played_games + 1 
		WHERE uuid=$1`,
		usersTable,
	))
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid)
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func (auth *AuthRepository) SetPlayedGames(numOfGames int, playerID int) error {
	const fn = "repository.postgres.auth.SetPlayedGames"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`UPDATE %s SET played_games=$1
		WHERE id=$2`,
		usersTable,
	))
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(numOfGames, playerID)
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func (auth *AuthRepository) IncrementConductedGames(playerID int) error {
	const fn = "repository.postgres.auth.IncrementConductedGames"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`UPDATE %s SET conducted_games = conducted_games + 1
		WHERE id=$1`,
		usersTable,
	))
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(playerID)
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func (auth *AuthRepository) SetConductedGames(numOfGames int, playerID int) error {
	const fn = "repository.postgres.auth.SetConductedGames"

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`UPDATE %s SET conducted_games = $1
		WHERE id=$2`,
		usersTable,
	))
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(numOfGames, playerID)
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func (auth *AuthRepository) SetStatus(status string, playerID int) error {
	const fn = "repository.postgres.auth.SetStatus"

	if status != playerStatus && status != masterStatus && status != adminStatus {
		return fmt.Errorf("%s: %v", fn, repository.ErrInvalidPlayerStatus)
	}

	stmt, err := auth.db.Prepare(fmt.Sprintf(
		`UPDATE %s SET status = $1
		WHERE id=$2`,
		usersTable,
	))
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, playerID)
	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func NewUUID() uuid.UUID {
	return uuid.New()
}
