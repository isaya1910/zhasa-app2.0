-- name: CreateUser :one
INSERT INTO users (phone, first_name, last_name)
VALUES ($1, $2, $3) ON CONFLICT (phone)
DO
UPDATE SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name
    RETURNING id;

-- name: GetUserByPhone :one
SELECT u.id,
       u.phone,
       u.first_name,
       u.last_name,
       u.avatar_url,
       u.about,
       u.work_start_date
FROM user_avatar_view u
WHERE u.phone = $1;

-- name: GetUserByPhoneWithPassword :one
SELECT u.id,
       u.phone,
       u.first_name,
       u.last_name,
       u.password
FROM users u
WHERE u.phone = $1;

-- name: GetUserById :one
SELECT *
FROM user_avatar_view u
         JOIN user_roles ur on u.id = ur.user_id
         JOIN roles r on ur.role_id = r.id
WHERE u.id = $1;

-- name: CreateUserCode :one
INSERT INTO users_codes(user_id, code)
VALUES ($1, $2) RETURNING id;

-- name: GetAuthCodeById :one
SELECT *
FROM users_codes
WHERE id = $1;

-- name: UploadUserAvatar :exec
INSERT INTO users_avatars(user_id, avatar_url)
VALUES ($1, $2) ON CONFLICT (user_id)
DO
UPDATE SET avatar_url = EXCLUDED.avatar_url;

-- name: DeleteUserAvatar :exec
DELETE
FROM users_avatars
WHERE user_id = $1;

-- name: GetUserBranch :one
SELECT b.title, b.id
FROM users u
         JOIN
     branch_users bu ON u.id = bu.user_id
         JOIN branches b ON bu.branch_id = b.id
WHERE u.id = $1;

-- name: SetUserBrandGoal :exec
INSERT INTO user_brand_sale_type_goals (user_id, brand_id, sale_type_id, value, from_date, to_date)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (user_id, brand_id, sale_type_id, from_date, to_date) DO
UPDATE
    SET value = $4;


-- name: GetUsersByBranchBrandRole :many
SELECT u.id,
       u.first_name,
       u.last_name,
       u.avatar_url
FROM user_avatar_view u
         JOIN user_brands ub ON u.id = ub.user_id AND ub.brand_id = $1
         JOIN branch_users bu ON u.id = bu.user_id AND bu.branch_id = $2
         JOIN user_roles ur ON u.id = ur.user_id AND ur.role_id = $3
         LEFT JOIN disabled_users du ON u.id = du.user_id
WHERE du.user_id IS NULL;

-- name: AddBrandToUser :exec
INSERT INTO user_brands (user_id, brand_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: AddRoleToUser :exec
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: AddUserToBranch :exec
INSERT INTO branch_users (user_id, branch_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: GetUsersWithoutRoles :many
SELECT u.id,
       u.first_name,
       u.last_name
FROM users u
         LEFT JOIN user_roles ur ON u.id = ur.user_id
WHERE ur.user_id IS NULL
  AND (u.last_name || ' ' || u.first_name) ILIKE '%' || @search::text || '%'
ORDER BY u.created_at DESC
    LIMIT 25;

-- name: GetUsersWithBranchRolesBrands :many
WITH Counted AS (SELECT u.id,
                        u.first_name,
                        u.last_name,
                        u.phone,
                        b.title                    AS branch_title,
                        STRING_AGG(bs.title, ', ') AS brands,
                        COUNT(*)                      OVER()            AS total_count
                 FROM users u
                          JOIN user_roles ur ON u.id = ur.user_id
                          JOIN roles r ON ur.role_id = r.id AND r.key = $1
                          JOIN branch_users bu ON u.id = bu.user_id
                          JOIN user_brands ub ON u.id = ub.user_id
                          JOIN brands bs ON ub.brand_id = bs.id
                          JOIN branches b ON bu.branch_id = b.id
                          LEFT JOIN disabled_users du ON u.id = du.user_id
                 WHERE (last_name || ' ' || first_name) ILIKE '%' || @search::text || '%'
GROUP BY u.id, u.first_name, u.last_name, b.title, du.user_id
    )
SELECT id,
       first_name,
       last_name,
       phone,
       branch_title,
       brands,
       total_count
FROM Counted
ORDER BY first_name, last_name, id DESC LIMIT $2
OFFSET $3;

-- name: UpdateUser :exec
UPDATE users
SET first_name = $1,
    last_name  = $2,
    phone      = $3
WHERE id = $4;

-- name: UpdateUserBranch :exec
UPDATE branch_users
SET branch_id = $1
WHERE user_id = $2;


-- name: GetUsersWithBranchBrands :many
WITH Counted AS (SELECT u.id,
                        u.first_name,
                        u.last_name,
                        u.phone,
                        b.title                    AS branch_title,
                        STRING_AGG(bs.title, ', ') AS brands,
                        COUNT(*)                      OVER()            AS total_count
                 FROM users u
                          JOIN branch_users bu ON u.id = bu.user_id
                          JOIN user_brands ub ON u.id = ub.user_id
                          JOIN brands bs ON ub.brand_id = bs.id
                          JOIN branches b ON bu.branch_id = b.id
                 WHERE (last_name || ' ' || first_name) ILIKE '%' || @search::text || '%'
GROUP BY u.id, u.first_name, u.last_name, b.title
    )
SELECT id,
       first_name,
       last_name,
       phone,
       branch_title,
       brands,
       total_count
FROM Counted
ORDER BY first_name, last_name, id DESC LIMIT $1
OFFSET $2;

-- name: AddDisabledUser :exec
INSERT INTO disabled_users (user_id)
VALUES ($1) ON CONFLICT DO NOTHING;

-- name: GetDisabledUser :one
SELECT user_id from disabled_users where user_id = $1;

-- name: GetFilteredUsersWithBranchRolesBrands :many
WITH Counted AS (
    SELECT u.id,
           u.first_name,
           u.last_name,
           u.phone,
           r.key                      AS role,
           b.title                    AS branch_title,
           STRING_AGG(bs.title, ', ') AS brands,
           COUNT(*) OVER()            AS total_count,
           (du.user_id IS NOT NULL)::boolean   AS deleted
    FROM users u
             JOIN user_roles ur ON u.id = ur.user_id
             JOIN roles r ON ur.role_id = r.id
             LEFT JOIN branch_users bu ON u.id = bu.user_id
             LEFT JOIN user_brands ub ON u.id = ub.user_id
             LEFT JOIN brands bs ON ub.brand_id = bs.id
             LEFT JOIN branches b ON bu.branch_id = b.id
             LEFT JOIN disabled_users du ON u.id = du.user_id
    WHERE (last_name || ' ' || first_name) ILIKE '%' || @search::text || '%'
      AND (@role_keys::text[] IS NULL OR r.key = ANY(@role_keys))
      AND (@brand_ids::int[] IS NULL OR bs.id = ANY(@brand_ids))
      AND (@branch_ids::int[] IS NULL OR b.id = ANY(@branch_ids))
      AND (@show_deleted::boolean OR du.user_id IS NULL)
    GROUP BY u.id, u.first_name, u.last_name, b.title, du.user_id, r.key
)
SELECT id,
       first_name,
       last_name,
       phone,
       role,
       branch_title,
       brands,
       total_count,
       deleted
FROM Counted
ORDER BY CASE WHEN @sort_field::text = 'fio' AND @sort_type::text = 'asc' THEN first_name END ASC,
    CASE WHEN @sort_field = 'fio' AND @sort_type = 'asc' THEN last_name
END ASC,
    CASE WHEN @sort_field = 'fio' AND @sort_type = 'desc' THEN first_name
END DESC,
    CASE WHEN @sort_field = 'fio' AND @sort_type = 'desc' THEN last_name
END DESC,
    CASE WHEN @sort_field = 'branch' AND @sort_type = 'asc' THEN branch_title
END ASC,
    CASE WHEN @sort_field = 'branch' AND @sort_type = 'desc' THEN branch_title
END DESC,
    first_name, last_name, id DESC
LIMIT $1 OFFSET $2;

-- name: AddUserRole :exec
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, (SELECT id FROM roles WHERE key = @role_key::text) )
ON CONFLICT DO NOTHING;

-- name: AddUserBranch :exec
INSERT INTO branch_users (user_id, branch_id)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: UpdateUserRole :exec
UPDATE user_roles
SET role_id = (SELECT id FROM roles WHERE key = @role_key::text)
WHERE user_id = $1;

-- name: DeleteUserBranchByUserId :exec
DELETE
FROM branch_users
WHERE user_id = $1;

-- name: DeleteDisabledUsers :exec
DELETE
FROM disabled_users
WHERE user_id = ANY($1::int[]);

-- name: SearchUsers :many
SELECT u.id,
       u.first_name,
       u.last_name,
       u.phone,
       u.avatar_url
FROM user_avatar_view u
JOIN user_roles ur ON u.id = ur.user_id
WHERE (u.last_name || ' ' || u.first_name) ILIKE '%' || $1 || '%'
LIMIT 20;

-- name: UpdateUserAbout :exec
UPDATE users
SET about = $1
WHERE id = $2;
