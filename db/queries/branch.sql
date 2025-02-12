-- name: GetBranchById :one
SELECT *
FROM branches
WHERE id = $1;

-- name: CreateBranch :exec
INSERT INTO branches (title, description)
VALUES ($1, $2);

-- name: GetBranchBrandGoalByGivenDateRange :one
SELECT COALESCE(bg.value, 0) AS goal_amount
FROM branch_brand_sale_type_goals bg
WHERE bg.branch_id = $1
  AND bg.brand_id = $2
  AND bg.from_date = $3
  AND bg.to_date = $4
  AND bg.sale_type_id = $5;

-- name: GetBranches :many
SELECT *
FROM branches;

-- name: GetBranchesByBrandId :many
SELECT b.id, b.title, b.description
FROM branches b
         JOIN branch_brands bb ON b.id = bb.branch_id
WHERE bb.brand_id = $1;

-- name: SetBranchBrandGoal :exec
INSERT INTO branch_brand_sale_type_goals (branch_id, brand_id, sale_type_id, value, from_date, to_date)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (branch_id, brand_id, sale_type_id, from_date, to_date) DO
UPDATE
    SET value = $4;

-- name: GetBranchBrandUserByRole :many
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
WHERE du.user_id IS NULL;

-- name: GetBranchBrandSaleSumByGivenDateRange :one
SELECT COALESCE(SUM(s.amount), 0) ::bigint AS total_sales
FROM sales s
         JOIN sales_brands sb ON s.id = sb.sale_id
         JOIN user_brands ub ON ub.user_id = s.user_id AND ub.brand_id = sb.brand_id
         JOIN branch_users bu ON bu.user_id = s.user_id
WHERE bu.branch_id = $1   -- Replace with the desired branch_id
  AND sb.brand_id = $2    -- Replace with the desired brand_id
  AND s.sale_type_id = $3 -- Replace with the desired sale_type_id
  AND s.sale_date BETWEEN $4 AND $5;
-- Replace with the desired period (from_date and to_date)

-- name: GetAllBranches :many
SELECT *
FROM branches;

-- name: SetBrandSaleTypeGoal :exec
INSERT INTO brand_overall_sale_type_goals (brand_id, sale_type_id, value, from_date, to_date)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (brand_id, sale_type_id, from_date, to_date) DO
UPDATE
    SET value = $3;

-- name: GetBrandSaleSumByGivenDateRange :one
SELECT COALESCE(SUM(s.amount), 0) ::bigint AS total_sales
FROM sales s
         JOIN sales_brands sb ON s.id = sb.sale_id
         JOIN user_brands ub ON ub.user_id = s.user_id AND ub.brand_id = sb.brand_id
WHERE sb.brand_id = $1    -- Replace with the desired brand_id
  AND s.sale_type_id = $2 -- Replace with the desired sale_type_id
  AND s.sale_date BETWEEN $3 AND $4;
-- Replace with the desired period (from_date and to_date)

-- name: GetBrandOverallGoalByGivenDateRange :one
SELECT COALESCE(bg.value, 0) AS goal_amount
FROM brand_overall_sale_type_goals bg
WHERE bg.brand_id = $1
  AND bg.from_date = $2
  AND bg.to_date = $3
  AND bg.sale_type_id = $4;

-- name: AddBranch :one
INSERT INTO branches (title, description)
VALUES ($1, $2)
RETURNING id;

-- name: AddBranchBrand :exec
INSERT INTO branch_brands (branch_id, brand_id)
VALUES ($1, $2);

-- name: UpdateBranch :exec
UPDATE branches
SET title = $1,
    description = $2
WHERE id = $3;

-- name: DeleteBranchBrands :exec
DELETE
FROM branch_brands
WHERE branch_id = $1;

-- name: GetBranchesSearchDesc :many
SELECT
    b.id,
    b.title,
    b.description,
    COALESCE(array_agg(DISTINCT br.title), '{}')::text AS brands
FROM
    branches b
        LEFT JOIN branch_brands bb ON bb.branch_id = b.id
        LEFT JOIN brands br ON br.id = bb.brand_id
WHERE
    (b.title || b.description) ILIKE '%' || @search::text || '%'
GROUP BY
    b.id, b.title, b.description
ORDER BY
    b.title DESC
    LIMIT $1 OFFSET $2;


-- name: GetBranchesSearchAsc :many
SELECT
    b.id,
    b.title,
    b.description,
    COALESCE(array_agg(DISTINCT br.title), '{}')::text AS brands
FROM
    branches b
        LEFT JOIN branch_brands bb ON bb.branch_id = b.id
        LEFT JOIN brands br ON br.id = bb.brand_id
WHERE
    (b.title || b.description) ILIKE '%' || @search::text || '%'
GROUP BY
    b.id, b.title, b.description
ORDER BY
    b.title ASC
    LIMIT $1 OFFSET $2;


-- name: GetBranchesSearchCount :one
select count(*)
from branches
where (title || description) ilike '%' || @search::text || '%';