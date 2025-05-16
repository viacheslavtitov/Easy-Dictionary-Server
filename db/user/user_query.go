package db

import (
	database "easy-dictionary-server/db"
	"fmt"
)

// GetAllUsersQuery get query to get all users from user table
func GetAllUsersQuery(orderBy database.OrderByType) string {
	return fmt.Sprintf(`SELECT * FROM user ORDER BY email %s`, orderBy.VALUE)
}

// GetUserByIdQuery get query to get user by id from user table
func GetUserByIdQuery() string {
	return `SELECT * FROM user WHERE id = $1`
}

// GetUserByEmailQuery get query to get user by email from user table
func GetUserByEmailQuery() string {
	return `SELECT * FROM user WHERE email = $1`
}

// CreateUserQuery get query to create user
// Params:
// - $1: email
// - $2: provider_id
// - $3: uid
// Return:
// - id: created user id
func CreateUserQuery() string {
	return `INSERT INTO user (email, provider_id, uid)
		VALUES ($1, $2, $3)
		RETURNING id`
}

// UpdateUserQuery get query to update user
// Params:
// - $1: email
// - $2: provider_id
// - $3: uid
// - $4: id of user in database which you want to update
func UpdateUserQuery() string {
	return `UPDATE users
		SET email = $1,
		    provider_id = $2,
		    uid = $3
		WHERE id = $4
		RETURNING id, email, provider_id, uid`
}

// DeleteUserByIdQuery get query to delete user by id from user table
// Params:
// - $1: id
func DeleteUserByIdQuery() string {
	return `DELETE FROM user WHERE id = $1`
}
