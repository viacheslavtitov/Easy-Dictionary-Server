package db

import (
	database "easy-dictionary-server/db"
	"fmt"
)

// GetAllUsersQuery get query to get all users from user table
func getAllUsersQuery(orderBy database.OrderByType) string {
	return fmt.Sprintf(`
SELECT 
    u.id AS user_id,
    u.uuid AS uuid,
    u.first_name,
	u.last_name,
    u.created_at AS user_created_at,
    u.user_role,
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
func getUserByIdQuery() string {
	return `
SELECT 
    u.id AS user_id,
    u.uuid AS uuid,
    u.first_name,
	u.last_name,
    u.created_at AS user_created_at,
    u.user_role,
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

// GetUserByUUIDQuery get query to get user by uuid from user table and join with user_providers table
func getUserByUUIDQuery() string {
	return `
SELECT 
    u.id AS user_id,
    u.uuid AS uuid,
    u.first_name,
	u.last_name,
    u.created_at AS user_created_at,
    u.user_role,
    p.id AS provider_id,
    p.provider_name,
	p.email,
    p.hashed_password,
    p.created_at AS provider_created_at
FROM users u
LEFT JOIN user_providers p
    ON u.id = p.user_id
WHERE u.uuid = $1;`
}

// GetUserByEmailQuery get query to get user by email from user table
func getUserByEmailQuery() string {
	return `
SELECT 
    u.id AS user_id,
    u.uuid AS uuid,
    u.first_name,
	u.last_name,
    u.created_at AS user_created_at,
    u.user_role,
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
// - $3: role
// - $4: provider name
// - $5: email
// - $6: hashed password
// Return:
// - uuid: created user uuid
func createUserQuery() string {
	return `
WITH new_user AS (
    INSERT INTO users (first_name, last_name, user_role, created_at)
    VALUES ($1, $2, $3, now())
    RETURNING id, uuid
)
inserted_provider AS (
    INSERT INTO user_providers (user_id, provider_name, email, hashed_password, created_at)
    VALUES (
        (SELECT id FROM new_user),
        $4, $5, $6, now()
    )
    RETURNING user_id
)
SELECT uuid FROM new_user;
`
}

// UpdateUserQuery get query to update user
// Params:
// - $1: first_name
// - $2: last_name
// - $3: id of user in database which you want to update
func updateUserQuery() string {
	return `
WITH updated_user AS (
    UPDATE users
    SET 
        first_name = $1,
        last_name = $2
    WHERE id = $3
    RETURNING id, first_name, last_name, created_at
)
SELECT 
    u.id AS user_id,
    u.first_name,
    u.last_name,
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
func deleteUserByIdQuery() string {
	return `DELETE FROM users WHERE id = $1`
}
