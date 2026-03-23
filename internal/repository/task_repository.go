package repository

import (
	"context"
	"database/sql"

	"awesomeProject1/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
	INSERT INTO tasks (type, payload, status)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at`

	return r.db.QueryRow(ctx, query,
		task.Type,
		task.Payload,
		task.Status,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetByID(ctx context.Context, id int64) (*model.Task, error) {
	query := `SELECT id, type, payload, status, result, error, created_at, updated_at FROM tasks WHERE id=$1`

	var task model.Task
	err := r.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.Type,
		&task.Payload,
		&task.Status,
		&task.Result,
		&task.Error,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	rows, err := r.db.Query(ctx, `SELECT id, type, payload, status, result, error, created_at, updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var t model.Task
		err := rows.Scan(
			&t.ID,
			&t.Type,
			&t.Payload,
			&t.Status,
			&t.Result,
			&t.Error,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) FetchPending(ctx context.Context) (*model.Task, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `
	SELECT id, type, payload FROM tasks
	WHERE status = 'pending'
	FOR UPDATE SKIP LOCKED
	LIMIT 1`

	var task model.Task

	err = tx.QueryRow(ctx, query).Scan(
		&task.ID,
		&task.Type,
		&task.Payload,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, `UPDATE tasks SET status='processing', updated_at=NOW() WHERE id=$1`, task.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	task.Status = model.StatusProcessing
	return &task, nil
}

func (r *TaskRepository) UpdateStatus(ctx context.Context, id int64, status model.TaskStatus, result, errMsg *string) error {
	_, err := r.db.Exec(ctx,
		`UPDATE tasks SET status=$1, result=$2, error=$3, updated_at=NOW() WHERE id=$4`,
		status, result, errMsg, id,
	)
	return err
}
