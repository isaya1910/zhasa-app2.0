// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: branch.sql

package generated

import (
	"context"
	"time"
)

const createBranch = `-- name: CreateBranch :exec
INSERT INTO branches (title, description)
VALUES ($1, $2)
`

type CreateBranchParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) CreateBranch(ctx context.Context, arg CreateBranchParams) error {
	_, err := q.db.ExecContext(ctx, createBranch, arg.Title, arg.Description)
	return err
}

const getAllBranches = `-- name: GetAllBranches :many

SELECT id, title, description, created_at
FROM branches
`

// Replace with the desired period (from_date and to_date)
func (q *Queries) GetAllBranches(ctx context.Context) ([]Branch, error) {
	rows, err := q.db.QueryContext(ctx, getAllBranches)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Branch
	for rows.Next() {
		var i Branch
		if err := rows.Scan(
			&i.ID,
			&i.Title,
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

const getBranchBrandGoalByGivenDateRange = `-- name: GetBranchBrandGoalByGivenDateRange :one
SELECT COALESCE(bg.value, 0) AS goal_amount
FROM branch_brand_sale_type_goals bg
WHERE bg.branch_id = $1
  AND bg.brand_id = $2
  AND bg.from_date = $3
  AND bg.to_date = $4
  AND bg.sale_type_id = $5
`

type GetBranchBrandGoalByGivenDateRangeParams struct {
	BranchID   int32     `json:"branch_id"`
	BrandID    int32     `json:"brand_id"`
	FromDate   time.Time `json:"from_date"`
	ToDate     time.Time `json:"to_date"`
	SaleTypeID int32     `json:"sale_type_id"`
}

func (q *Queries) GetBranchBrandGoalByGivenDateRange(ctx context.Context, arg GetBranchBrandGoalByGivenDateRangeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBranchBrandGoalByGivenDateRange,
		arg.BranchID,
		arg.BrandID,
		arg.FromDate,
		arg.ToDate,
		arg.SaleTypeID,
	)
	var goal_amount int64
	err := row.Scan(&goal_amount)
	return goal_amount, err
}

const getBranchBrandSaleSumByGivenDateRange = `-- name: GetBranchBrandSaleSumByGivenDateRange :one
SELECT COALESCE(SUM(s.amount), 0) ::bigint AS total_sales
FROM sales s
         JOIN sales_brands sb ON s.id = sb.sale_id
         JOIN user_brands ub ON ub.user_id = s.user_id AND ub.brand_id = sb.brand_id
         JOIN branch_users bu ON bu.user_id = s.user_id
WHERE bu.branch_id = $1   -- Replace with the desired branch_id
  AND sb.brand_id = $2    -- Replace with the desired brand_id
  AND s.sale_type_id = $3 -- Replace with the desired sale_type_id
  AND s.sale_date BETWEEN $4 AND $5
`

type GetBranchBrandSaleSumByGivenDateRangeParams struct {
	BranchID   int32     `json:"branch_id"`
	BrandID    int32     `json:"brand_id"`
	SaleTypeID int32     `json:"sale_type_id"`
	SaleDate   time.Time `json:"sale_date"`
	SaleDate_2 time.Time `json:"sale_date_2"`
}

func (q *Queries) GetBranchBrandSaleSumByGivenDateRange(ctx context.Context, arg GetBranchBrandSaleSumByGivenDateRangeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBranchBrandSaleSumByGivenDateRange,
		arg.BranchID,
		arg.BrandID,
		arg.SaleTypeID,
		arg.SaleDate,
		arg.SaleDate_2,
	)
	var total_sales int64
	err := row.Scan(&total_sales)
	return total_sales, err
}

const getBranchBrandUserByRole = `-- name: GetBranchBrandUserByRole :many
SELECT u.id,
       u.first_name,
       u.last_name,
       u.avatar_url,
       b.title AS branch_title,
       b.id    AS branch_id
FROM user_avatar_view u
         JOIN user_brands ub ON u.id = ub.user_id AND ub.brand_id = $1
         JOIN branch_users bu ON u.id = bu.user_id AND bu.branch_id = $2
         JOIN branches b ON bu.branch_id = b.id
         JOIN user_roles ur ON u.id = ur.user_id AND ur.role_id = $3
         LEFT JOIN disabled_users du ON u.id = du.user_id
WHERE du.user_id IS NULL
`

type GetBranchBrandUserByRoleParams struct {
	BrandID  int32 `json:"brand_id"`
	BranchID int32 `json:"branch_id"`
	RoleID   int32 `json:"role_id"`
}

type GetBranchBrandUserByRoleRow struct {
	ID          int32  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AvatarUrl   string `json:"avatar_url"`
	BranchTitle string `json:"branch_title"`
	BranchID    int32  `json:"branch_id"`
}

func (q *Queries) GetBranchBrandUserByRole(ctx context.Context, arg GetBranchBrandUserByRoleParams) ([]GetBranchBrandUserByRoleRow, error) {
	rows, err := q.db.QueryContext(ctx, getBranchBrandUserByRole, arg.BrandID, arg.BranchID, arg.RoleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBranchBrandUserByRoleRow
	for rows.Next() {
		var i GetBranchBrandUserByRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.AvatarUrl,
			&i.BranchTitle,
			&i.BranchID,
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

const getBranchById = `-- name: GetBranchById :one
SELECT id, title, description, created_at
FROM branches
WHERE id = $1
`

func (q *Queries) GetBranchById(ctx context.Context, id int32) (Branch, error) {
	row := q.db.QueryRowContext(ctx, getBranchById, id)
	var i Branch
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getBranches = `-- name: GetBranches :many
SELECT id, title, description, created_at
FROM branches
`

func (q *Queries) GetBranches(ctx context.Context) ([]Branch, error) {
	rows, err := q.db.QueryContext(ctx, getBranches)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Branch
	for rows.Next() {
		var i Branch
		if err := rows.Scan(
			&i.ID,
			&i.Title,
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

const getBranchesByBrandId = `-- name: GetBranchesByBrandId :many
SELECT b.id, b.title, b.description
FROM branches b
         JOIN branch_brands bb ON b.id = bb.branch_id
WHERE bb.brand_id = $1
`

type GetBranchesByBrandIdRow struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) GetBranchesByBrandId(ctx context.Context, brandID int32) ([]GetBranchesByBrandIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getBranchesByBrandId, brandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBranchesByBrandIdRow
	for rows.Next() {
		var i GetBranchesByBrandIdRow
		if err := rows.Scan(&i.ID, &i.Title, &i.Description); err != nil {
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

const getBrandOverallGoalByGivenDateRange = `-- name: GetBrandOverallGoalByGivenDateRange :one

SELECT COALESCE(bg.value, 0) AS goal_amount
FROM brand_overall_sale_type_goals bg
WHERE bg.brand_id = $1
  AND bg.from_date = $2
  AND bg.to_date = $3
  AND bg.sale_type_id = $4
`

type GetBrandOverallGoalByGivenDateRangeParams struct {
	BrandID    int32     `json:"brand_id"`
	FromDate   time.Time `json:"from_date"`
	ToDate     time.Time `json:"to_date"`
	SaleTypeID int32     `json:"sale_type_id"`
}

// Replace with the desired period (from_date and to_date)
func (q *Queries) GetBrandOverallGoalByGivenDateRange(ctx context.Context, arg GetBrandOverallGoalByGivenDateRangeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBrandOverallGoalByGivenDateRange,
		arg.BrandID,
		arg.FromDate,
		arg.ToDate,
		arg.SaleTypeID,
	)
	var goal_amount int64
	err := row.Scan(&goal_amount)
	return goal_amount, err
}

const getBrandSaleSumByGivenDateRange = `-- name: GetBrandSaleSumByGivenDateRange :one
SELECT COALESCE(SUM(s.amount), 0) ::bigint AS total_sales
FROM sales s
         JOIN sales_brands sb ON s.id = sb.sale_id
         JOIN user_brands ub ON ub.user_id = s.user_id AND ub.brand_id = sb.brand_id
WHERE sb.brand_id = $1    -- Replace with the desired brand_id
  AND s.sale_type_id = $2 -- Replace with the desired sale_type_id
  AND s.sale_date BETWEEN $3 AND $4
`

type GetBrandSaleSumByGivenDateRangeParams struct {
	BrandID    int32     `json:"brand_id"`
	SaleTypeID int32     `json:"sale_type_id"`
	SaleDate   time.Time `json:"sale_date"`
	SaleDate_2 time.Time `json:"sale_date_2"`
}

func (q *Queries) GetBrandSaleSumByGivenDateRange(ctx context.Context, arg GetBrandSaleSumByGivenDateRangeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBrandSaleSumByGivenDateRange,
		arg.BrandID,
		arg.SaleTypeID,
		arg.SaleDate,
		arg.SaleDate_2,
	)
	var total_sales int64
	err := row.Scan(&total_sales)
	return total_sales, err
}

const setBranchBrandGoal = `-- name: SetBranchBrandGoal :exec
INSERT INTO branch_brand_sale_type_goals (branch_id, brand_id, sale_type_id, value, from_date, to_date)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (branch_id, brand_id, sale_type_id, from_date, to_date) DO
UPDATE
    SET value = $4
`

type SetBranchBrandGoalParams struct {
	BranchID   int32     `json:"branch_id"`
	BrandID    int32     `json:"brand_id"`
	SaleTypeID int32     `json:"sale_type_id"`
	Value      int64     `json:"value"`
	FromDate   time.Time `json:"from_date"`
	ToDate     time.Time `json:"to_date"`
}

func (q *Queries) SetBranchBrandGoal(ctx context.Context, arg SetBranchBrandGoalParams) error {
	_, err := q.db.ExecContext(ctx, setBranchBrandGoal,
		arg.BranchID,
		arg.BrandID,
		arg.SaleTypeID,
		arg.Value,
		arg.FromDate,
		arg.ToDate,
	)
	return err
}

const setBrandSaleTypeGoal = `-- name: SetBrandSaleTypeGoal :exec
INSERT INTO brand_overall_sale_type_goals (brand_id, sale_type_id, value, from_date, to_date)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (brand_id, sale_type_id, from_date, to_date) DO
UPDATE
    SET value = $3
`

type SetBrandSaleTypeGoalParams struct {
	BrandID    int32     `json:"brand_id"`
	SaleTypeID int32     `json:"sale_type_id"`
	Value      int64     `json:"value"`
	FromDate   time.Time `json:"from_date"`
	ToDate     time.Time `json:"to_date"`
}

func (q *Queries) SetBrandSaleTypeGoal(ctx context.Context, arg SetBrandSaleTypeGoalParams) error {
	_, err := q.db.ExecContext(ctx, setBrandSaleTypeGoal,
		arg.BrandID,
		arg.SaleTypeID,
		arg.Value,
		arg.FromDate,
		arg.ToDate,
	)
	return err
}
