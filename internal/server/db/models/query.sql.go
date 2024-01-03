// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const addPortToConnection = `-- name: AddPortToConnection :exec
UPDATE connections
SET
    port = ?
WHERE
    id = ?
`

type AddPortToConnectionParams struct {
	Port interface{}
	ID   string
}

func (q *Queries) AddPortToConnection(ctx context.Context, arg AddPortToConnectionParams) error {
	_, err := q.db.ExecContext(ctx, addPortToConnection, arg.Port, arg.ID)
	return err
}

const createGlobalSettings = `-- name: CreateGlobalSettings :one
INSERT INTO
    global_settings (
        smtp_enabled,
        add_member_email_subject,
        add_member_email_template
    )
VALUES
    (?, ?, ?) RETURNING id, smtp_enabled, smtp_host, smtp_port, smtp_username, smtp_password, from_address, add_member_email_subject, add_member_email_template
`

type CreateGlobalSettingsParams struct {
	SmtpEnabled            bool
	AddMemberEmailSubject  interface{}
	AddMemberEmailTemplate interface{}
}

func (q *Queries) CreateGlobalSettings(ctx context.Context, arg CreateGlobalSettingsParams) (GlobalSetting, error) {
	row := q.db.QueryRowContext(ctx, createGlobalSettings, arg.SmtpEnabled, arg.AddMemberEmailSubject, arg.AddMemberEmailTemplate)
	var i GlobalSetting
	err := row.Scan(
		&i.ID,
		&i.SmtpEnabled,
		&i.SmtpHost,
		&i.SmtpPort,
		&i.SmtpUsername,
		&i.SmtpPassword,
		&i.FromAddress,
		&i.AddMemberEmailSubject,
		&i.AddMemberEmailTemplate,
	)
	return i, err
}

const createNewHttpConnection = `-- name: CreateNewHttpConnection :one
INSERT INTO
    connections (id, type, subdomain, team_member_id, team_id)
VALUES
    (?, "http", ?, ?, ?) RETURNING id, type, subdomain, port, status, team_member_id, created_at, started_at, closed_at, team_id
`

type CreateNewHttpConnectionParams struct {
	ID           string
	Subdomain    interface{}
	TeamMemberID int64
	TeamID       interface{}
}

func (q *Queries) CreateNewHttpConnection(ctx context.Context, arg CreateNewHttpConnectionParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, createNewHttpConnection,
		arg.ID,
		arg.Subdomain,
		arg.TeamMemberID,
		arg.TeamID,
	)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Subdomain,
		&i.Port,
		&i.Status,
		&i.TeamMemberID,
		&i.CreatedAt,
		&i.StartedAt,
		&i.ClosedAt,
		&i.TeamID,
	)
	return i, err
}

const createNewTcpConnection = `-- name: CreateNewTcpConnection :one
INSERT INTO
    connections (id, type, port, team_member_id, team_id)
VALUES
    (?, "tcp", ?, ?, ?) RETURNING id, type, subdomain, port, status, team_member_id, created_at, started_at, closed_at, team_id
`

type CreateNewTcpConnectionParams struct {
	ID           string
	Port         interface{}
	TeamMemberID int64
	TeamID       interface{}
}

func (q *Queries) CreateNewTcpConnection(ctx context.Context, arg CreateNewTcpConnectionParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, createNewTcpConnection,
		arg.ID,
		arg.Port,
		arg.TeamMemberID,
		arg.TeamID,
	)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Subdomain,
		&i.Port,
		&i.Status,
		&i.TeamMemberID,
		&i.CreatedAt,
		&i.StartedAt,
		&i.ClosedAt,
		&i.TeamID,
	)
	return i, err
}

const createSession = `-- name: CreateSession :one
INSERT INTO
    sessions (token, user_id)
VALUES
    (?, ?) RETURNING id, user_id, token, created_at
`

type CreateSessionParams struct {
	Token  string
	UserID int64
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.Token, arg.UserID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.CreatedAt,
	)
	return i, err
}

const createTeam = `-- name: CreateTeam :one
INSERT INTO
    teams (name, slug)
VALUES
    (?, ?) RETURNING id, name, slug, created_at
`

type CreateTeamParams struct {
	Name string
	Slug string
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRowContext(ctx, createTeam, arg.Name, arg.Slug)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreatedAt,
	)
	return i, err
}

const createTeamMember = `-- name: CreateTeamMember :one
INSERT INTO
    team_members (user_id, team_id, role, secret_key)
VALUES
    (?, ?, ?, ?) RETURNING id, user_id, team_id, secret_key, role, added_by_user_id, created_at
`

type CreateTeamMemberParams struct {
	UserID    int64
	TeamID    int64
	Role      string
	SecretKey string
}

func (q *Queries) CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, createTeamMember,
		arg.UserID,
		arg.TeamID,
		arg.Role,
		arg.SecretKey,
	)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (
        email,
        first_name,
        last_name,
        is_super_user,
        github_access_token,
        github_avatar_url
    )
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING id, email, first_name, last_name, is_super_user, github_access_token, github_avatar_url, created_at
`

type CreateUserParams struct {
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.IsSuperUser,
		arg.GithubAccessToken,
		arg.GithubAvatarUrl,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExpiredSessions = `-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE
    strftime ('%s', 'now') - strftime ('%s', created_at) > 24 * 60 * 60
`

func (q *Queries) DeleteExpiredSessions(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExpiredSessions)
	return err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE
    token = ?
`

func (q *Queries) DeleteSession(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteSession, token)
	return err
}

const deleteUnclaimedConnections = `-- name: DeleteUnclaimedConnections :exec
DELETE FROM connections
WHERE
    status = 'reserved'
    AND strftime ('%s', 'now') - strftime ('%s', created_at) > 10
`

func (q *Queries) DeleteUnclaimedConnections(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUnclaimedConnections)
	return err
}

const getActiveConnectionsForTeam = `-- name: GetActiveConnectionsForTeam :many
SELECT
    connections.id,
    connections.type,
    connections.port,
    connections.subdomain,
    connections.created_at,
    connections.started_at,
    connections.closed_at,
    connections.status,
    users.email,
    users.first_name,
    users.last_name,
    users.github_avatar_url
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
    JOIN users ON users.id = team_members.user_id
WHERE
    connections.team_id = ?
    AND status = 'active'
ORDER BY
    connections.id DESC
LIMIT
    20
`

type GetActiveConnectionsForTeamRow struct {
	ID              string
	Type            string
	Port            interface{}
	Subdomain       interface{}
	CreatedAt       time.Time
	StartedAt       interface{}
	ClosedAt        interface{}
	Status          string
	Email           string
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
}

func (q *Queries) GetActiveConnectionsForTeam(ctx context.Context, teamID interface{}) ([]GetActiveConnectionsForTeamRow, error) {
	rows, err := q.db.QueryContext(ctx, getActiveConnectionsForTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetActiveConnectionsForTeamRow
	for rows.Next() {
		var i GetActiveConnectionsForTeamRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Port,
			&i.Subdomain,
			&i.CreatedAt,
			&i.StartedAt,
			&i.ClosedAt,
			&i.Status,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.GithubAvatarUrl,
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

const getGlobalSettings = `-- name: GetGlobalSettings :one
SELECT
    id, smtp_enabled, smtp_host, smtp_port, smtp_username, smtp_password, from_address, add_member_email_subject, add_member_email_template
FROM
    global_settings
LIMIT
    1
`

func (q *Queries) GetGlobalSettings(ctx context.Context) (GlobalSetting, error) {
	row := q.db.QueryRowContext(ctx, getGlobalSettings)
	var i GlobalSetting
	err := row.Scan(
		&i.ID,
		&i.SmtpEnabled,
		&i.SmtpHost,
		&i.SmtpPort,
		&i.SmtpUsername,
		&i.SmtpPassword,
		&i.FromAddress,
		&i.AddMemberEmailSubject,
		&i.AddMemberEmailTemplate,
	)
	return i, err
}

const getRecentConnectionsForTeam = `-- name: GetRecentConnectionsForTeam :many
SELECT
    connections.id,
    connections.type,
    connections.port,
    connections.subdomain,
    connections.created_at,
    connections.started_at,
    connections.closed_at,
    connections.status,
    users.email,
    users.first_name,
    users.last_name,
    users.github_avatar_url
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
    JOIN users ON users.id = team_members.user_id
WHERE
    connections.team_id = ?
    AND status != 'reserved'
ORDER BY
    connections.id DESC
LIMIT
    20
`

type GetRecentConnectionsForTeamRow struct {
	ID              string
	Type            string
	Port            interface{}
	Subdomain       interface{}
	CreatedAt       time.Time
	StartedAt       interface{}
	ClosedAt        interface{}
	Status          string
	Email           string
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
}

func (q *Queries) GetRecentConnectionsForTeam(ctx context.Context, teamID interface{}) ([]GetRecentConnectionsForTeamRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentConnectionsForTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentConnectionsForTeamRow
	for rows.Next() {
		var i GetRecentConnectionsForTeamRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Port,
			&i.Subdomain,
			&i.CreatedAt,
			&i.StartedAt,
			&i.ClosedAt,
			&i.Status,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.GithubAvatarUrl,
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

const getReservedOrActiveConnectionById = `-- name: GetReservedOrActiveConnectionById :one
SELECT
    connections.id, type, subdomain, port, status, team_member_id, connections.created_at, started_at, closed_at, connections.team_id, team_members.id, user_id, team_members.team_id, secret_key, role, added_by_user_id, team_members.created_at
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
WHERE
    connections.id = ?
    AND status IN ('active', 'reserved')
LIMIT
    1
`

type GetReservedOrActiveConnectionByIdRow struct {
	ID            string
	Type          string
	Subdomain     interface{}
	Port          interface{}
	Status        string
	TeamMemberID  int64
	CreatedAt     time.Time
	StartedAt     interface{}
	ClosedAt      interface{}
	TeamID        int64
	ID_2          int64
	UserID        int64
	TeamID_2      int64
	SecretKey     string
	Role          string
	AddedByUserID interface{}
	CreatedAt_2   time.Time
}

func (q *Queries) GetReservedOrActiveConnectionById(ctx context.Context, id string) (GetReservedOrActiveConnectionByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getReservedOrActiveConnectionById, id)
	var i GetReservedOrActiveConnectionByIdRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Subdomain,
		&i.Port,
		&i.Status,
		&i.TeamMemberID,
		&i.CreatedAt,
		&i.StartedAt,
		&i.ClosedAt,
		&i.TeamID,
		&i.ID_2,
		&i.UserID,
		&i.TeamID_2,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt_2,
	)
	return i, err
}

const getReservedOrActiveConnectionForPort = `-- name: GetReservedOrActiveConnectionForPort :one
SELECT
    connections.id, type, subdomain, port, status, team_member_id, connections.created_at, started_at, closed_at, connections.team_id, team_members.id, user_id, team_members.team_id, secret_key, role, added_by_user_id, team_members.created_at
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
WHERE
    port = ?
    AND team_members.secret_key = ?
    AND status IN ('active', 'reserved')
LIMIT
    1
`

type GetReservedOrActiveConnectionForPortParams struct {
	Port      interface{}
	SecretKey string
}

type GetReservedOrActiveConnectionForPortRow struct {
	ID            string
	Type          string
	Subdomain     interface{}
	Port          interface{}
	Status        string
	TeamMemberID  int64
	CreatedAt     time.Time
	StartedAt     interface{}
	ClosedAt      interface{}
	TeamID        int64
	ID_2          int64
	UserID        int64
	TeamID_2      int64
	SecretKey     string
	Role          string
	AddedByUserID interface{}
	CreatedAt_2   time.Time
}

func (q *Queries) GetReservedOrActiveConnectionForPort(ctx context.Context, arg GetReservedOrActiveConnectionForPortParams) (GetReservedOrActiveConnectionForPortRow, error) {
	row := q.db.QueryRowContext(ctx, getReservedOrActiveConnectionForPort, arg.Port, arg.SecretKey)
	var i GetReservedOrActiveConnectionForPortRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Subdomain,
		&i.Port,
		&i.Status,
		&i.TeamMemberID,
		&i.CreatedAt,
		&i.StartedAt,
		&i.ClosedAt,
		&i.TeamID,
		&i.ID_2,
		&i.UserID,
		&i.TeamID_2,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt_2,
	)
	return i, err
}

const getReservedOrActiveConnectionForSubdomain = `-- name: GetReservedOrActiveConnectionForSubdomain :one
SELECT
    connections.id, type, subdomain, port, status, team_member_id, connections.created_at, started_at, closed_at, connections.team_id, team_members.id, user_id, team_members.team_id, secret_key, role, added_by_user_id, team_members.created_at
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
WHERE
    subdomain = ?
    AND team_members.secret_key = ?
    AND status IN ('active', 'reserved')
LIMIT
    1
`

type GetReservedOrActiveConnectionForSubdomainParams struct {
	Subdomain interface{}
	SecretKey string
}

type GetReservedOrActiveConnectionForSubdomainRow struct {
	ID            string
	Type          string
	Subdomain     interface{}
	Port          interface{}
	Status        string
	TeamMemberID  int64
	CreatedAt     time.Time
	StartedAt     interface{}
	ClosedAt      interface{}
	TeamID        int64
	ID_2          int64
	UserID        int64
	TeamID_2      int64
	SecretKey     string
	Role          string
	AddedByUserID interface{}
	CreatedAt_2   time.Time
}

func (q *Queries) GetReservedOrActiveConnectionForSubdomain(ctx context.Context, arg GetReservedOrActiveConnectionForSubdomainParams) (GetReservedOrActiveConnectionForSubdomainRow, error) {
	row := q.db.QueryRowContext(ctx, getReservedOrActiveConnectionForSubdomain, arg.Subdomain, arg.SecretKey)
	var i GetReservedOrActiveConnectionForSubdomainRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Subdomain,
		&i.Port,
		&i.Status,
		&i.TeamMemberID,
		&i.CreatedAt,
		&i.StartedAt,
		&i.ClosedAt,
		&i.TeamID,
		&i.ID_2,
		&i.UserID,
		&i.TeamID_2,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamById = `-- name: GetTeamById :one
SELECT
    id, name, slug, created_at
FROM
    teams
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetTeamById(ctx context.Context, id int64) (Team, error) {
	row := q.db.QueryRowContext(ctx, getTeamById, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreatedAt,
	)
	return i, err
}

const getTeamMemberByEmail = `-- name: GetTeamMemberByEmail :one
SELECT
    team_members.id, user_id, team_id, secret_key, role, added_by_user_id, team_members.created_at, users.id, email, first_name, last_name, is_super_user, github_access_token, github_avatar_url, users.created_at
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
WHERE
    users.email = ?
LIMIT
    1
`

type GetTeamMemberByEmailRow struct {
	ID                int64
	UserID            int64
	TeamID            int64
	SecretKey         string
	Role              string
	AddedByUserID     interface{}
	CreatedAt         time.Time
	ID_2              int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt_2       time.Time
}

func (q *Queries) GetTeamMemberByEmail(ctx context.Context, email string) (GetTeamMemberByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberByEmail, email)
	var i GetTeamMemberByEmailRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamMemberById = `-- name: GetTeamMemberById :one
SELECT
    team_members.id, team_members.user_id, team_members.team_id, team_members.secret_key, team_members.role, team_members.added_by_user_id, team_members.created_at,
    users.id, users.email, users.first_name, users.last_name, users.is_super_user, users.github_access_token, users.github_avatar_url, users.created_at
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
WHERE
    team_members.id = ?
LIMIT
    1
`

type GetTeamMemberByIdRow struct {
	ID                int64
	UserID            int64
	TeamID            int64
	SecretKey         string
	Role              string
	AddedByUserID     interface{}
	CreatedAt         time.Time
	ID_2              int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt_2       time.Time
}

func (q *Queries) GetTeamMemberById(ctx context.Context, id int64) (GetTeamMemberByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberById, id)
	var i GetTeamMemberByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamMemberByUserIdAndTeamSlug = `-- name: GetTeamMemberByUserIdAndTeamSlug :one
SELECT
    team_members.id, team_members.user_id, team_members.team_id, team_members.secret_key, team_members.role, team_members.added_by_user_id, team_members.created_at,
    users.id, users.email, users.first_name, users.last_name, users.is_super_user, users.github_access_token, users.github_avatar_url, users.created_at
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
    JOIN teams ON teams.id = team_members.team_id
WHERE
    users.id = ?
    AND teams.slug = ?
LIMIT
    1
`

type GetTeamMemberByUserIdAndTeamSlugParams struct {
	ID   int64
	Slug string
}

type GetTeamMemberByUserIdAndTeamSlugRow struct {
	ID                int64
	UserID            int64
	TeamID            int64
	SecretKey         string
	Role              string
	AddedByUserID     interface{}
	CreatedAt         time.Time
	ID_2              int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt_2       time.Time
}

func (q *Queries) GetTeamMemberByUserIdAndTeamSlug(ctx context.Context, arg GetTeamMemberByUserIdAndTeamSlugParams) (GetTeamMemberByUserIdAndTeamSlugRow, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberByUserIdAndTeamSlug, arg.ID, arg.Slug)
	var i GetTeamMemberByUserIdAndTeamSlugRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamMembers = `-- name: GetTeamMembers :many
SELECT
    users.email,
    team_members.role,
    users.github_avatar_url
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
WHERE
    team_id = ?
`

type GetTeamMembersRow struct {
	Email           string
	Role            string
	GithubAvatarUrl interface{}
}

func (q *Queries) GetTeamMembers(ctx context.Context, teamID int64) ([]GetTeamMembersRow, error) {
	rows, err := q.db.QueryContext(ctx, getTeamMembers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTeamMembersRow
	for rows.Next() {
		var i GetTeamMembersRow
		if err := rows.Scan(&i.Email, &i.Role, &i.GithubAvatarUrl); err != nil {
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

const getTeamUserBySecretKey = `-- name: GetTeamUserBySecretKey :one
SELECT
    id, user_id, team_id, secret_key, role, added_by_user_id, created_at
FROM
    team_members
WHERE
    secret_key = ?
LIMIT
    1
`

func (q *Queries) GetTeamUserBySecretKey(ctx context.Context, secretKey string) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, getTeamUserBySecretKey, secretKey)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
	)
	return i, err
}

const getTeamsOfUser = `-- name: GetTeamsOfUser :many
SELECT
    teams.id, teams.name, teams.slug, teams.created_at
FROM
    team_members
    JOIN teams ON teams.id = team_members.team_id
WHERE
    team_members.user_id = ?
`

func (q *Queries) GetTeamsOfUser(ctx context.Context, userID int64) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getTeamsOfUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.CreatedAt,
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    id, email, first_name, last_name, is_super_user, github_access_token, github_avatar_url, created_at
FROM
    users
WHERE
    email = ?
LIMIT
    1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT
    users.id,
    users.email,
    users.created_at,
    users.first_name,
    users.last_name,
    users.github_avatar_url,
    users.is_super_user
FROM
    users
WHERE
    id = ?
LIMIT
    1
`

type GetUserByIdRow struct {
	ID              int64
	Email           string
	CreatedAt       time.Time
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
	IsSuperUser     bool
}

func (q *Queries) GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.GithubAvatarUrl,
		&i.IsSuperUser,
	)
	return i, err
}

const getUserBySession = `-- name: GetUserBySession :one
SELECT
    users.id,
    users.email,
    users.created_at,
    users.first_name,
    users.last_name,
    users.github_avatar_url,
    users.is_super_user
FROM
    users
    JOIN sessions ON sessions.user_id = users.id
WHERE
    sessions.token = ?
LIMIT
    1
`

type GetUserBySessionRow struct {
	ID              int64
	Email           string
	CreatedAt       time.Time
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
	IsSuperUser     bool
}

func (q *Queries) GetUserBySession(ctx context.Context, token string) (GetUserBySessionRow, error) {
	row := q.db.QueryRowContext(ctx, getUserBySession, token)
	var i GetUserBySessionRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.GithubAvatarUrl,
		&i.IsSuperUser,
	)
	return i, err
}

const getUsersCount = `-- name: GetUsersCount :one
SELECT
    COUNT(*)
FROM
    users
`

func (q *Queries) GetUsersCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUsersCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const markConnectionAsActive = `-- name: MarkConnectionAsActive :exec
UPDATE connections
SET
    status = 'active',
    started_at = CURRENT_TIMESTAMP
WHERE
    id = ?
`

func (q *Queries) MarkConnectionAsActive(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, markConnectionAsActive, id)
	return err
}

const markConnectionAsClosed = `-- name: MarkConnectionAsClosed :exec
UPDATE connections
SET
    status = 'closed',
    closed_at = CURRENT_TIMESTAMP
WHERE
    id = ?
`

func (q *Queries) MarkConnectionAsClosed(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, markConnectionAsClosed, id)
	return err
}

const updateGlobalSettings = `-- name: UpdateGlobalSettings :exec
UPDATE global_settings
SET
    smtp_enabled = ?,
    smtp_host = ?,
    smtp_port = ?,
    smtp_username = ?,
    smtp_password = ?,
    from_address = ?,
    add_member_email_subject = ?,
    add_member_email_template = ?
`

type UpdateGlobalSettingsParams struct {
	SmtpEnabled            bool
	SmtpHost               interface{}
	SmtpPort               interface{}
	SmtpUsername           interface{}
	SmtpPassword           interface{}
	FromAddress            interface{}
	AddMemberEmailSubject  interface{}
	AddMemberEmailTemplate interface{}
}

func (q *Queries) UpdateGlobalSettings(ctx context.Context, arg UpdateGlobalSettingsParams) error {
	_, err := q.db.ExecContext(ctx, updateGlobalSettings,
		arg.SmtpEnabled,
		arg.SmtpHost,
		arg.SmtpPort,
		arg.SmtpUsername,
		arg.SmtpPassword,
		arg.FromAddress,
		arg.AddMemberEmailSubject,
		arg.AddMemberEmailTemplate,
	)
	return err
}

const updateSecretKey = `-- name: UpdateSecretKey :exec
UPDATE team_members
SET
    secret_key = ?
WHERE
    id = ?
`

type UpdateSecretKeyParams struct {
	SecretKey string
	ID        int64
}

func (q *Queries) UpdateSecretKey(ctx context.Context, arg UpdateSecretKeyParams) error {
	_, err := q.db.ExecContext(ctx, updateSecretKey, arg.SecretKey, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
    first_name = COALESCE(?, first_name),
    last_name = COALESCE(?, last_name),
    github_access_token = COALESCE(?, github_access_token),
    github_avatar_url = COALESCE(?, github_avatar_url)
WHERE
    id = ?
`

type UpdateUserParams struct {
	FirstName         interface{}
	LastName          interface{}
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	ID                int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.GithubAccessToken,
		arg.GithubAvatarUrl,
		arg.ID,
	)
	return err
}
