// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID      `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  []byte           `json:"password"`
	Role      string           `json:"role"`
	Verified  pgtype.Bool      `json:"verified"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
