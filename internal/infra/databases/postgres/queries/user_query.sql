-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (id_user, name, last_name, email) 
VALUES ($1,$2,$3,$4)
RETURNING *;