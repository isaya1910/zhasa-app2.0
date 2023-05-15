// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: branch.sql

package generated

import (
	"context"
	"time"
)

const createBranch = `-- name: CreateBranch :exec
INSERT INTO branches (title, description, branch_key)
VALUES ($1, $2, $3)
`

type CreateBranchParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	BranchKey   string `json:"branch_key"`
}

func (q *Queries) CreateBranch(ctx context.Context, arg CreateBranchParams) error {
	_, err := q.db.ExecContext(ctx, createBranch, arg.Title, arg.Description, arg.BranchKey)
	return err
}

const getBranchById = `-- name: GetBranchById :one
SELECT id, title, description, branch_key, created_at FROM branches
WHERE id = $1
`

func (q *Queries) GetBranchById(ctx context.Context, id int32) (Branch, error) {
	row := q.db.QueryRowContext(ctx, getBranchById, id)
	var i Branch
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.BranchKey,
		&i.CreatedAt,
	)
	return i, err
}

const getBranchesByRating = `-- name: GetBranchesByRating :many
WITH sales_summary AS (
    SELECT
        b.id AS branch_id,
        SUM(s.amount) AS total_sales_amount
    FROM
        sales s
            INNER JOIN sales_managers sm ON s.sales_manager_id = sm.id
            INNER JOIN branches b ON sm.branch_id = b.id
    WHERE
        s.sale_date BETWEEN $1 AND $2
    GROUP BY
        b.id
),
     goal_summary AS (
         SELECT
             b.id AS branch_id,
             SUM(smg.amount) AS total_goal_amount
         FROM
             sales_manager_goals smg
                 INNER JOIN sales_managers sm ON smg.sales_manager_id = sm.id
                 INNER JOIN branches b ON sm.branch_id = b.id
         WHERE
                 smg.from_date = $1
           AND smg.to_date = $2
         GROUP BY
             b.id
     )
SELECT
    b.id AS branch_id,
    b.title AS branch_title,
    b.branch_key AS branch_key,
    b.description AS description,
    COALESCE(ss.total_sales_amount / NULLIF(smg.total_goal_amount, 0), 0)::float AS ratio
FROM
    branches b
        LEFT JOIN sales_summary ss ON b.id = ss.branch_id
        LEFT JOIN goal_summary smg ON b.id = smg.branch_id
ORDER BY
    ratio DESC
`

type GetBranchesByRatingParams struct {
	SaleDate   time.Time `json:"sale_date"`
	SaleDate_2 time.Time `json:"sale_date_2"`
}

type GetBranchesByRatingRow struct {
	BranchID    int32   `json:"branch_id"`
	BranchTitle string  `json:"branch_title"`
	BranchKey   string  `json:"branch_key"`
	Description string  `json:"description"`
	Ratio       float64 `json:"ratio"`
}

// Get Ranked Branches
func (q *Queries) GetBranchesByRating(ctx context.Context, arg GetBranchesByRatingParams) ([]GetBranchesByRatingRow, error) {
	rows, err := q.db.QueryContext(ctx, getBranchesByRating, arg.SaleDate, arg.SaleDate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBranchesByRatingRow
	for rows.Next() {
		var i GetBranchesByRatingRow
		if err := rows.Scan(
			&i.BranchID,
			&i.BranchTitle,
			&i.BranchKey,
			&i.Description,
			&i.Ratio,
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