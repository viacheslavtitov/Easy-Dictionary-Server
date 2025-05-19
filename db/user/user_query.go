package db

import (
	database "easy-dictionary-server/db"
	"fmt"
)

// GetAllUsersQuery get query to get all users from user table
func GetAllUsersQuery(orderBy database.OrderByType) string {
	return fmt.Sprintf(`
SELECT 
    u.id AS user_id,
    u.first_name,
	u.second_name,
    u.created_at AS user_created_at,
    p.id AS provider_id,
    p.provider_name,
	p.email,
    p.hashed_password,
    p.created_at AS provider_created_at
FROM users u
LEFT JOIN user_providers p
    ON u.id = p.user_id
ORDER BY p.email %s;`, orderBy.VALUE)
}

// GetUserByIdQuery get query to get user by id from user table and join with user_providers table
func GetUserByIdQuery() string {
	return `
SELECT 
    u.id AS user_id,
    u.first_name,
	u.second_name,
    u.created_at AS user_created_at,
    p.id AS provider_id,
    p.provider_name,
	p.email,
    p.hashed_password,
    p.created_at AS provider_created_at
FROM users u
LEFT JOIN user_providers p
    ON u.id = p.user_id
WHERE u.id = $1;`
}

// GetUserByEmailQuery get query to get user by email from user table
func GetUserByEmailQuery() string {
	return `
SELECT 
    u.id AS user_id,
    u.first_name,
	u.second_name,
    u.created_at AS user_created_at,
    p.id AS provider_id,
    p.provider_name,
	p.email,
    p.hashed_password,
    p.created_at AS provider_created_at
FROM users u
LEFT JOIN user_providers p
    ON u.id = p.user_id
WHERE p.email = $1;`
}

// CreateUserQuery get query to create user
// Params:
// - $1: first name
// - $2: second name
// - $3: provider name
// - $4: email
// - $5: hashed password
// Return:
// - id: created user id
func CreateUserQuery() string {
	return `
WITH new_user AS (
    INSERT INTO users (first_name, second_name, created_at)
    VALUES ($1, $2, now())
    RETURNING id
)
INSERT INTO user_providers (user_id, provider_name, email, hashed_password, created_at)
VALUES (
    (SELECT id FROM new_user),
    $3, $4, $5, now()
)
RETURNING *;
`
}

// UpdateUserQuery get query to update user
// Params:
// - $1: first_name
// - $2: second_name
// - $3: id of user in database which you want to update
func UpdateUserQuery() string {
	return `
WITH updated_user AS (
    UPDATE users
    SET 
        first_name = $1,
        second_name = $2
    WHERE id = $3
    RETURNING id, first_name, second_name, created_at
)
SELECT 
    u.id AS user_id,
    u.first_name,
    u.second_name,
    u.created_at AS user_created_at,

    p.id AS provider_id,
    p.provider_name,
	p.email,
    p.hashed_password,
    p.created_at AS provider_created_at
FROM updated_user u
LEFT JOIN user_providers p ON u.id = p.user_id;`
}

// DeleteUserByIdQuery get query to delete user by id from user table
// Params:
// - $1: id
func DeleteUserByIdQuery() string {
	return `DELETE FROM user WHERE id = $1`
}
