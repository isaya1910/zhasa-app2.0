// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: branch_director.sql

package generated

import (
	"context"
	"time"
)

const createBranchDirector = `-- name: CreateBranchDirector :one
INSERT INTO branch_directors (user_id, branch_id)
VALUES ($1, $2) RETURNING id
`

type CreateBranchDirectorParams struct {
	UserID   int32 `json:"user_id"`
	BranchID int32 `json:"branch_id"`
}

func (q *Queries) CreateBranchDirector(ctx context.Context, arg CreateBranchDirectorParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createBranchDirector, arg.UserID, arg.BranchID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createSalesManagerGoal = `-- name: CreateSalesManagerGoal :exec
INSERT INTO sales_manager_goals (sales_manager_id, from_date, to_date, amount)
VALUES ($1, $2, $3, $4)
`

type CreateSalesManagerGoalParams struct {
	SalesManagerID int32     `json:"sales_manager_id"`
	FromDate       time.Time `json:"from_date"`
	ToDate         time.Time `json:"to_date"`
	Amount         int64     `json:"amount"`
}

func (q *Queries) CreateSalesManagerGoal(ctx context.Context, arg CreateSalesManagerGoalParams) error {
	_, err := q.db.ExecContext(ctx, createSalesManagerGoal,
		arg.SalesManagerID,
		arg.FromDate,
		arg.ToDate,
		arg.Amount,
	)
	return err
}

const getBranchDirectorByUserId = `-- name: GetBranchDirectorByUserId :one
SELECT user_id, phone, first_name, last_name, avatar_url, branch_director_id, branch_id, branch_title FROM branch_directors_view bdv
WHERE bdv.user_id = $1
`

func (q *Queries) GetBranchDirectorByUserId(ctx context.Context, userID int32) (BranchDirectorsView, error) {
	row := q.db.QueryRowContext(ctx, getBranchDirectorByUserId, userID)
	var i BranchDirectorsView
	err := row.Scan(
		&i.UserID,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.AvatarUrl,
		&i.BranchDirectorID,
		&i.BranchID,
		&i.BranchTitle,
	)
	return i, err
}
