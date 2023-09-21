package auth

import (
	"errors"
	"log"
	"time"

	db "github.com/yogapratama23/go-http-server/internal/database"
)

type AuthRepository struct{}

func (r *AuthRepository) Create(p *SignupInput) error {
	params := []interface{}{p.Username, p.Password, time.Now(), time.Now()}

	_, err := db.Connect.Exec(`
		INSERT INTO 
			users (
				username, password, created_at, updated_at
			)
		VALUES (
			?, ?, ?, ?
		)
	`, params...)
	if err != nil {
		log.Println(err)
		return errors.New("signup failed")
	}
	return nil
}

func (r *AuthRepository) FindByUsername(u *string) (*UserResponse, error) {
	var user UserResponse
	params := []interface{}{u}
	query := `
		SELECT
			id, username, password
		FROM
			users
		WHERE
			deleted_at IS NULL
		AND
			username = ?
	`

	err := db.Connect.QueryRow(query, params...).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *AuthRepository) FindById(id *int) (*UserResponse, error) {
	var user UserResponse
	params := []interface{}{id}
	query := `
		SELECT
			id, username, password
		FROM
			users
		WHERE
			deleted_at IS NULL
		AND
			id = ?
	`

	err := db.Connect.QueryRow(query, params...).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *AuthRepository) FindToken(t string) (*TokenResponse, error) {
	var token TokenResponse
	params := []interface{}{t}
	query := `
		SELECT
			token,
			user_id
		FROM
			tokens
		WHERE
			token = ?
		AND
			deleted_at IS NULL
	`

	err := db.Connect.QueryRow(query, params...).Scan(&token.Token, &token.UserId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("token not found")
	}

	return &token, nil
}

func (r *AuthRepository) CreateToken(id *int, token *string) error {
	params := []interface{}{token, id, time.Now(), time.Now()}
	query := `
		INSERT INTO
			tokens (
				token,
				user_id,
				created_at,
				updated_at
			)
		VALUES (
			?, ?, ?, ?
		)
	`

	_, err := db.Connect.Exec(query, params...)
	if err != nil {
		log.Println(err)
		return errors.New("signin failed")
	}

	return nil
}

func (r *AuthRepository) SoftDeleteToken(p *SignoutInput) error {
	params := []interface{}{time.Now(), p.Token, p.ID}
	query := `
		UPDATE
			tokens
		SET
			deleted_at = ?
		WHERE
			token = ?
		AND
			user_id = ?
	`

	_, err := db.Connect.Exec(query, params...)
	if err != nil {
		log.Println(err)
		return errors.New("delete token failed")
	}

	return nil
}
