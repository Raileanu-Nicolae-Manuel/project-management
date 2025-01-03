// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: project_users.sql

package db

import (
	"context"
	"database/sql"
)

const addUserToProject = `-- name: AddUserToProject :exec
INSERT INTO user_projects (
    user_id, project_id, role_id
) VALUES (
    ?, ?, ?
)
`

type AddUserToProjectParams struct {
	UserID    int64
	ProjectID int64
	RoleID    int64
}

func (q *Queries) AddUserToProject(ctx context.Context, arg AddUserToProjectParams) error {
	_, err := q.db.ExecContext(ctx, addUserToProject, arg.UserID, arg.ProjectID, arg.RoleID)
	return err
}

const getProjectUsers = `-- name: GetProjectUsers :many
SELECT u.id, u.username, u.email, u.password_hash, u.created_at, u.updated_at, pr.name as role_name
FROM users u
JOIN user_projects up ON u.id = up.user_id
JOIN project_roles pr ON up.role_id = pr.id
WHERE up.project_id = ?
`

type GetProjectUsersRow struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	RoleName     string
}

func (q *Queries) GetProjectUsers(ctx context.Context, projectID int64) ([]GetProjectUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getProjectUsers, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectUsersRow
	for rows.Next() {
		var i GetProjectUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.PasswordHash,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RoleName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserProjects = `-- name: GetUserProjects :many
SELECT p.id, p.name, p.description, p.created_at, p.updated_at, pr.name as role_name
FROM projects p
JOIN user_projects up ON p.id = up.project_id
JOIN project_roles pr ON up.role_id = pr.id
WHERE up.user_id = ?
`

type GetUserProjectsRow struct {
	ID          int64
	Name        string
	Description sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	RoleName    string
}

func (q *Queries) GetUserProjects(ctx context.Context, userID int64) ([]GetUserProjectsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserProjects, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserProjectsRow
	for rows.Next() {
		var i GetUserProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RoleName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeUserFromProject = `-- name: RemoveUserFromProject :exec
DELETE FROM user_projects
WHERE user_id = ? AND project_id = ?
`

type RemoveUserFromProjectParams struct {
	UserID    int64
	ProjectID int64
}

func (q *Queries) RemoveUserFromProject(ctx context.Context, arg RemoveUserFromProjectParams) error {
	_, err := q.db.ExecContext(ctx, removeUserFromProject, arg.UserID, arg.ProjectID)
	return err
}

const updateUserProjectRole = `-- name: UpdateUserProjectRole :exec
UPDATE user_projects
SET role_id = ?
WHERE user_id = ? AND project_id = ?
`

type UpdateUserProjectRoleParams struct {
	RoleID    int64
	UserID    int64
	ProjectID int64
}

func (q *Queries) UpdateUserProjectRole(ctx context.Context, arg UpdateUserProjectRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateUserProjectRole, arg.RoleID, arg.UserID, arg.ProjectID)
	return err
}
