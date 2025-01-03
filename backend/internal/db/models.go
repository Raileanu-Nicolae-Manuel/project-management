// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type Project struct {
	ID          int64
	Name        string
	Description sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

type ProjectRole struct {
	ID          int64
	Name        string
	Description sql.NullString
	CreatedAt   sql.NullTime
}

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

type UserProject struct {
	UserID    int64
	ProjectID int64
	CreatedAt sql.NullTime
	RoleID    int64
}

type UserRole struct {
	ID          int32
	Name        string
	Description sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

type UserRoleAssignment struct {
	UserID    sql.NullInt64
	RoleID    sql.NullInt32
	CreatedAt sql.NullTime
}
