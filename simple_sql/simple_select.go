package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(
	ctx context.Context,
	conn *pgx.Conn,
) ([]TasksModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]TasksModel, 0)

	for rows.Next() {
		var task TasksModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		fmt.Println(task)
	}
	return tasks, nil
}
