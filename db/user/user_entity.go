package db

import (
	database "easy-dictionary-server/db"
)

type UserEntity struct {
	ID         int    `db:"id"`
	Email      string `db:"email"`
	ProviderID string `db:"provider_id"`
	UID        string `db:"uid"`
}

func GetAllUsers(db *database.Database, orderBy database.OrderByType) ([]UserEntity, error) {
	users := []UserEntity{}
	err := db.SQLDB.Select(&users, GetAllUsersQuery(orderBy))
	return users, err
}

func GetUserById(db *database.Database, id int) (UserEntity, error) {
	user := UserEntity{}
	err := db.SQLDB.Select(&user, GetUserByIdQuery(), id)
	return user, err
}

func GetUserByEmail(db *database.Database, email string) (UserEntity, error) {
	user := UserEntity{}
	err := db.SQLDB.Select(&user, GetUserByEmailQuery(), email)
	return user, err
}

func CreateUser(db *database.Database, user UserEntity) (int, error) {
	createdId := -1
	err := db.SQLDB.Select(&createdId, CreateUserQuery(), user.Email, user.ProviderID, user.UID)
	return createdId, err
}

func UpdateUser(db *database.Database, user UserEntity) error {
	_, err := db.SQLDB.NamedExec(UpdateUserQuery(), user)
	return err
}

func DeleteUserById(db *database.Database, id int) error {
	_, err := db.SQLDB.Exec(DeleteUserByIdQuery(), id)
	return err
}
