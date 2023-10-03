package repositories

import (
	"context"
	"database/sql"
	"sqap/internal/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func todoFromRows(row sql.Row) (*models.Todo, error) {
	var t models.Todo
	err := row.Scan(&t.UID, &t.Content, &t.IsComplete)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *TodoRepository) Get(ctx context.Context) (*models.Todo, error) {
	query := "SELECT (content, is_complete) FROM todos LIMIT 1;"

	row := r.db.QueryRowContext(ctx, query)

	todo, err := todoFromRows(*row)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) error {
	query := "INSERT INTO todos (uid, content, is_complete) VALUES (?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	is_complete := 0
	if todo.IsComplete {
		is_complete = 1
	}

	_, err = stmt.ExecContext(
		ctx,
		todo.UID,
		todo.Content,
		is_complete,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) GetTodo(ctx context.Context, uid string) (*models.Todo, error) {
	query := "SELECT (content, is_complete) FROM todos WHERE uid = ?"

	row := r.db.QueryRowContext(ctx, query, uid)

	todo, err := todoFromRows(*row)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
