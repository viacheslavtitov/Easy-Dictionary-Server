package db

import (
	database "easy-dictionary-server/db"
	pointers "easy-dictionary-server/internalenv/utils"
	"errors"
	"time"
)

type UserEntity struct {
	ID        int                   `db:"id"`
	FirstName string                `db:"first_name"`
	LastName  string                `db:"last_name"`
	Role      string                `db:"user_role"`
	Providers *[]UserProviderEntity `db:"-"`
	CreatedAt time.Time             `db:"created_at"`
}

type UserProviderEntity struct {
	ID             int       `db:"id"`
	UserId         int       `db:"user_id"`
	ProviderName   string    `db:"provider_name"`
	HashedPassword string    `db:"hashed_password"`
	Email          string    `db:"email"`
	CreatedAt      time.Time `db:"created_at"`
}

type userWithProviderRow struct {
	UserID          int        `db:"user_id"`
	FirstName       string     `db:"first_name"`
	LastName        string     `db:"last_name"`
	UserCreatedAt   time.Time  `db:"user_created_at"`
	Role            string     `db:"user_role"`
	ProviderID      *int       `db:"provider_id"`
	ProviderName    *string    `db:"provider_name"`
	Email           *string    `db:"email"`
	HashedPassword  *string    `db:"hashed_password"`
	ProviderCreated *time.Time `db:"provider_created_at"`
}

func GetAllUsers(db *database.Database, orderBy database.OrderByType) ([]UserEntity, error) {
	var rows []userWithProviderRow
	err := db.SQLDB.Select(&rows, getAllUsersQuery(orderBy))
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return []UserEntity{}, nil
	}
	userMap := make(map[int]*UserEntity)
	for _, row := range rows {
		user, exists := userMap[row.UserID]
		if !exists {
			user = &UserEntity{
				ID:        row.UserID,
				FirstName: row.FirstName,
				LastName:  row.LastName,
				CreatedAt: row.UserCreatedAt,
				Providers: &[]UserProviderEntity{},
			}
			userMap[row.UserID] = user
		}

		if row.ProviderID != nil {
			*user.Providers = append(*user.Providers, UserProviderEntity{
				ID:             *row.ProviderID,
				UserId:         row.UserID,
				ProviderName:   pointers.Deref(row.ProviderName),
				Email:          pointers.Deref(row.Email),
				HashedPassword: pointers.Deref(row.HashedPassword),
				CreatedAt:      *row.ProviderCreated,
			})
		}
	}
	users := make([]UserEntity, 0, len(userMap))
	for _, u := range userMap {
		users = append(users, *u)
	}
	return users, err
}

func GetUserById(db *database.Database, id int) (*UserEntity, error) {
	var rows []userWithProviderRow
	err := db.SQLDB.Select(&rows, getUserByIdQuery(), id)
	return mapUserWithProvidersToEntity(err, rows)
}

func GetUserByEmail(db *database.Database, email string) (*UserEntity, error) {
	var rows []userWithProviderRow
	err := db.SQLDB.Select(&rows, getUserByEmailQuery(), email)
	return mapUserWithProvidersToEntity(err, rows)
}

func mapUserWithProvidersToEntity(err error, rows []userWithProviderRow) (*UserEntity, error) {
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errors.New("User not found")
	}
	user := UserEntity{
		ID:        rows[0].UserID,
		FirstName: rows[0].FirstName,
		LastName:  rows[0].LastName,
		CreatedAt: rows[0].UserCreatedAt,
		Providers: &[]UserProviderEntity{},
	}
	for _, row := range rows {
		if row.ProviderID != nil {
			*user.Providers = append(*user.Providers, UserProviderEntity{
				ID:             *row.ProviderID,
				UserId:         row.UserID,
				ProviderName:   pointers.Deref(row.ProviderName),
				Email:          pointers.Deref(row.Email),
				HashedPassword: pointers.Deref(row.HashedPassword),
			})
		}
	}
	return &user, err
}

func CreateUser(db *database.Database, user *UserEntity) (int, error) {
	createdId := -1
	err := db.SQLDB.Get(&createdId, createUserQuery(), user.FirstName, user.LastName, user.Role,
		(*user.Providers)[0].ProviderName, (*user.Providers)[0].Email, (*user.Providers)[0].HashedPassword)
	if err != nil {
		return -1, err
	}
	return createdId, nil
}

func UpdateUser(db *database.Database, user UserEntity) error {
	_, err := db.SQLDB.Exec(updateUserQuery(), user.FirstName, user.LastName, user.ID)
	return err
}

func DeleteUserById(db *database.Database, id int) error {
	_, err := db.SQLDB.Exec(deleteUserByIdQuery(), id)
	return err
}
