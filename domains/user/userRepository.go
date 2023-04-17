package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email, role) VALUES($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email, user.Role).
		Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password, role FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).
		Scan(&u.ID, &u.Email, &u.Username, &u.Password, &u.Role)

	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}

func (r *repository) GetUserById(ctx context.Context, id int64) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password, role FROM users WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&u.ID, &u.Email, &u.Username, &u.Password, &u.Role)

	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}

func (q *repository) GetUsers(ctx context.Context) ([]User, error) {
	query := "SELECT id, email, username, password, role FROM users"
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.Role,
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
