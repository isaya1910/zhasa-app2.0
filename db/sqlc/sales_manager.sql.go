// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: sales_manager.sql

package generated

import (
	"context"
	"time"
)

const addSaleOrReplace = `-- name: AddSaleOrReplace :exec
INSERT INTO sales (sales_manager_id, sale_date, amount, sale_type_id, description)
VALUES ($1, $2, $3, $4, $5)
`

type AddSaleOrReplaceParams struct {
	SalesManagerID int32     `json:"sales_manager_id"`
	SaleDate       time.Time `json:"sale_date"`
	Amount         int64     `json:"amount"`
	SaleTypeID     int32     `json:"sale_type_id"`
	Description    string    `json:"description"`
}

// add sale into sales by given sale_type_id, amount, date, sales_manager_id and on conflict replace
func (q *Queries) AddSaleOrReplace(ctx context.Context, arg AddSaleOrReplaceParams) error {
	_, err := q.db.ExecContext(ctx, addSaleOrReplace,
		arg.SalesManagerID,
		arg.SaleDate,
		arg.Amount,
		arg.SaleTypeID,
		arg.Description,
	)
	return err
}

const createSalesManager = `-- name: CreateSalesManager :exec
INSERT INTO sales_managers (user_id, branch_id)
VALUES ($1, $2)
`

type CreateSalesManagerParams struct {
	UserID   int32 `json:"user_id"`
	BranchID int32 `json:"branch_id"`
}

func (q *Queries) CreateSalesManager(ctx context.Context, arg CreateSalesManagerParams) error {
	_, err := q.db.ExecContext(ctx, createSalesManager, arg.UserID, arg.BranchID)
	return err
}

const getManagerSales = `-- name: GetManagerSales :many
SELECT id, sale_type_id, description, sale_date, amount
FROM sales s
WHERE s.sales_manager_id = $1
ORDER BY s.sale_date
LIMIT $2
OFFSET $3
`

type GetManagerSalesParams struct {
	SalesManagerID int32 `json:"sales_manager_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

type GetManagerSalesRow struct {
	ID          int32     `json:"id"`
	SaleTypeID  int32     `json:"sale_type_id"`
	Description string    `json:"description"`
	SaleDate    time.Time `json:"sale_date"`
	Amount      int64     `json:"amount"`
}

func (q *Queries) GetManagerSales(ctx context.Context, arg GetManagerSalesParams) ([]GetManagerSalesRow, error) {
	rows, err := q.db.QueryContext(ctx, getManagerSales, arg.SalesManagerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetManagerSalesRow
	for rows.Next() {
		var i GetManagerSalesRow
		if err := rows.Scan(
			&i.ID,
			&i.SaleTypeID,
			&i.Description,
			&i.SaleDate,
			&i.Amount,
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

const getRankedSalesManagers = `-- name: GetRankedSalesManagers :many
WITH sales_summary AS (SELECT sm.id         AS sales_manager_id,
                              SUM(s.amount) AS total_sales_amount,
                              u.first_name  AS first_name,
                              u.last_name   AS last_name,
                              u.avatar_url  AS avatar_url
                       FROM sales s
                                INNER JOIN sales_managers sm ON s.sales_manager_id = sm.id
                                INNER JOIN user_avatar_view u ON sm.user_id = u.id
                       WHERE s.sale_date BETWEEN $1 AND $2
                       GROUP BY sm.id),
     goal_summary AS (SELECT sm.id     AS sales_manager_id,
                             sg.from_date,
                             sg.to_date,
                             sg.amount AS goal_amount
                      FROM sales_manager_goals sg
                               INNER JOIN sales_managers sm ON s.sales_manager_id = sm.id
                      WHERE sg.from_date = $1
                        AND sg.to_date = $2)
SELECT ss.sales_manager_id,
       ss.first_name,
       ss.last_name,
       ss.avatar_url,
       COALESCE(ss.total_sales_amount / NULLIF(smg.goal_amount, 0), 0) ::float AS ratio
FROM sales_summary ss
         LEFT JOIN goal_summary smg ON ss.sales_manager_id = smg.sales_manager_id
ORDER BY ratio DESC LIMIT $3
OFFSET $4
`

type GetRankedSalesManagersParams struct {
	SaleDate   time.Time `json:"sale_date"`
	SaleDate_2 time.Time `json:"sale_date_2"`
	Limit      int32     `json:"limit"`
	Offset     int32     `json:"offset"`
}

type GetRankedSalesManagersRow struct {
	SalesManagerID int32   `json:"sales_manager_id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	AvatarUrl      string  `json:"avatar_url"`
	Ratio          float64 `json:"ratio"`
}

// get the ranked sales managers by their total sales divided by their sales goal amount for the given period.
func (q *Queries) GetRankedSalesManagers(ctx context.Context, arg GetRankedSalesManagersParams) ([]GetRankedSalesManagersRow, error) {
	rows, err := q.db.QueryContext(ctx, getRankedSalesManagers,
		arg.SaleDate,
		arg.SaleDate_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRankedSalesManagersRow
	for rows.Next() {
		var i GetRankedSalesManagersRow
		if err := rows.Scan(
			&i.SalesManagerID,
			&i.FirstName,
			&i.LastName,
			&i.AvatarUrl,
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

const getSalesByDate = `-- name: GetSalesByDate :many
SELECT id, sales_manager_id, sale_date, amount, sale_type_id, description, created_at
from sales s
WHERE s.sale_date = $1
`

func (q *Queries) GetSalesByDate(ctx context.Context, saleDate time.Time) ([]Sale, error) {
	rows, err := q.db.QueryContext(ctx, getSalesByDate, saleDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sale
	for rows.Next() {
		var i Sale
		if err := rows.Scan(
			&i.ID,
			&i.SalesManagerID,
			&i.SaleDate,
			&i.Amount,
			&i.SaleTypeID,
			&i.Description,
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

const getSalesManagerByUserId = `-- name: GetSalesManagerByUserId :one
SELECT user_id, phone, first_name, last_name, avatar_url, sales_manager_id, branch_id, branch_title
from sales_managers_view s
WHERE s.user_id = $1
`

func (q *Queries) GetSalesManagerByUserId(ctx context.Context, userID int32) (SalesManagersView, error) {
	row := q.db.QueryRowContext(ctx, getSalesManagerByUserId, userID)
	var i SalesManagersView
	err := row.Scan(
		&i.UserID,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.AvatarUrl,
		&i.SalesManagerID,
		&i.BranchID,
		&i.BranchTitle,
	)
	return i, err
}

const getSalesManagerGoalByGivenDateRange = `-- name: GetSalesManagerGoalByGivenDateRange :one
SELECT COALESCE(sg.amount, 0) AS goal_amount
FROM sales_manager_goals sg
WHERE sg.sales_manager_id = $1
  AND sg.from_date = $2
  AND sg.to_date = $3
`

type GetSalesManagerGoalByGivenDateRangeParams struct {
	SalesManagerID int32     `json:"sales_manager_id"`
	FromDate       time.Time `json:"from_date"`
	ToDate         time.Time `json:"to_date"`
}

func (q *Queries) GetSalesManagerGoalByGivenDateRange(ctx context.Context, arg GetSalesManagerGoalByGivenDateRangeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getSalesManagerGoalByGivenDateRange, arg.SalesManagerID, arg.FromDate, arg.ToDate)
	var goal_amount int64
	err := row.Scan(&goal_amount)
	return goal_amount, err
}

const getSalesManagerSumsByType = `-- name: GetSalesManagerSumsByType :many
SELECT st.id         AS sale_type_id,
       st.title      AS sale_type_title,
       SUM(s.amount) AS total_sales
FROM sale_types st
         JOIN sales s ON st.id = s.sale_type_id AND s.sales_manager_id = $1 AND s.sale_date BETWEEN $2 AND $3
GROUP BY st.id
ORDER BY st.id ASC
`

type GetSalesManagerSumsByTypeParams struct {
	SalesManagerID int32     `json:"sales_manager_id"`
	SaleDate       time.Time `json:"sale_date"`
	SaleDate_2     time.Time `json:"sale_date_2"`
}

type GetSalesManagerSumsByTypeRow struct {
	SaleTypeID    int32  `json:"sale_type_id"`
	SaleTypeTitle string `json:"sale_type_title"`
	TotalSales    int64  `json:"total_sales"`
}

// get the sales sums for a specific sales manager and each sale type within the given period.
func (q *Queries) GetSalesManagerSumsByType(ctx context.Context, arg GetSalesManagerSumsByTypeParams) ([]GetSalesManagerSumsByTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, getSalesManagerSumsByType, arg.SalesManagerID, arg.SaleDate, arg.SaleDate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSalesManagerSumsByTypeRow
	for rows.Next() {
		var i GetSalesManagerSumsByTypeRow
		if err := rows.Scan(&i.SaleTypeID, &i.SaleTypeTitle, &i.TotalSales); err != nil {
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
