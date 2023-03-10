package server

import (
	"database/sql"
	"errors"
	"github.com/deepmap/oapi-codegen/examples/petstore-expanded/echo/api/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserServer struct {
	DB *sql.DB
}

func NewUserServer(db *sql.DB) *UserServer {
	return &UserServer{DB: db}
}

func (s UserServer) CreateUser(ctx echo.Context) error {
	var u User

	err := ctx.Bind(&u)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for User")
	}

	query := `
        INSERT INTO users (email, first_name, last_name, phone, username) 
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

	args := []any{u.Email, u.FirstName, u.LastName, u.Phone, u.Username}

	err = s.DB.QueryRow(query, args...).Scan(&u.Id)
	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusCreated, u)
	if err != nil {
		return err
	}

	return nil
}

func (s UserServer) DeleteUser(ctx echo.Context, userId int64) error {
	if userId < 1 {
		return sendError(ctx, http.StatusNotFound, "User not found")
	}

	query := `
        DELETE FROM users
        WHERE id = $1`

	result, err := s.DB.Exec(query, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sendError(ctx, http.StatusNotFound, "User not found")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s UserServer) FindUserById(ctx echo.Context, userId int64) error {
	if userId < 1 {
		return sendError(ctx, http.StatusNotFound, "User not found")
	}

	query := `
        SELECT id, email, first_name, last_name, phone, username
        FROM users
        WHERE id = $1`

	var u User

	err := s.DB.QueryRow(query, userId).Scan(
		&u.Id,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.Phone,
		&u.Username,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return sendError(ctx, http.StatusNotFound, "User not found")
		default:
			return sendError(ctx, http.StatusInternalServerError, err.Error())
		}
	}

	err = ctx.JSON(http.StatusOK, u)
	if err != nil {
		return err
	}

	return nil
}

func (s UserServer) UpdateUser(ctx echo.Context, userId int64) error {
	var u User

	if userId < 1 {
		return sendError(ctx, http.StatusNotFound, "User not found")
	}

	query := `
        SELECT id, email, first_name, last_name, phone, username
        FROM users
        WHERE id = $1`

	err := s.DB.QueryRow(query, userId).Scan(
		&u.Id,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.Phone,
		&u.Username,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return sendError(ctx, http.StatusNotFound, "User not found")
		default:
			return sendError(ctx, http.StatusInternalServerError, err.Error())
		}
	}

	err = ctx.Bind(&u)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for User")
	}

	// Update
	query = `
        UPDATE users 
        SET email = $1, first_name = $2, last_name = $3, phone = $4, username = $5
        WHERE id = $6`

	args := []any{
		u.Email,
		u.FirstName,
		u.LastName,
		u.Phone,
		u.Username,
		userId,
	}

	err = s.DB.QueryRow(query, args...).Err()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return sendError(ctx, http.StatusNotFound, "User not found")
		default:
			return sendError(ctx, http.StatusInternalServerError, err.Error())
		}
	}

	err = ctx.JSON(http.StatusOK, u)
	if err != nil {
		return err
	}

	return nil
}

func sendError(ctx echo.Context, code int, message string) error {
	err := models.Error{
		Code:    int32(code),
		Message: message,
	}

	return ctx.JSON(code, err)
}
