package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"

	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

func Timeout(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
}

type Author struct {
	Id   string
	Name string
}
type AuthorMethods interface {
	Create(ctx context.Context, user Author) error
	FindAll(ctx context.Context) ([]Author, error)
	FindById(ctx context.Context, id string) (Author, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user Author) error
}

func (c *Connection) Create(ctx context.Context, user Author) (string, error) {
	Timeout(ctx)
	var id string
	q := `INSERT INTO author (name) VALUES ($1) RETURNING id;`
	if err := c.pool.QueryRow(ctx, q, user.Name).Scan(&id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return "", fmt.Errorf(pgErr.Message)
		}
		return "", fmt.Errorf("failed to insert author: %w", err)
	}
	return "author created", nil
}

func (c *Connection) FindAll(ctx context.Context) ([]Author, error) {
	Timeout(ctx)
	q := `SELECT id, name FROM author`
	rows, err := c.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to query authors: %w", err)
	}
	defer rows.Close()
	var authors []Author
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.Id, &author.Name); err != nil {
			return nil, fmt.Errorf("failed to scan author: %w", err)
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (c *Connection) FindById(ctx context.Context, id string) (Author, error) {
	Timeout(ctx)
	q := `SELECT id, name FROM author WHERE id = $1`
	var author Author
	if err := c.pool.QueryRow(ctx, q, id).Scan(&author.Id, &author.Name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return author, fmt.Errorf("author not found")
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return author, fmt.Errorf("pgx error: %w", pgErr)
		}
		return author, fmt.Errorf("failed to query author: %w", err)
	}
	return author, nil
}

func (c *Connection) Delete(ctx context.Context, id string) error {
	Timeout(ctx)

	q := `DELETE FROM author WHERE id = $1`
	pgTag, err := c.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}
	if pgTag.RowsAffected() == 0 {
		return fmt.Errorf("author not found")
	}
	fmt.Println("Deleted author:", id)
	return nil
}

func (c *Connection) Update(ctx context.Context, user Author) error {
	//TODO implement me
	panic("лень бля")
}
