package db

import (
	database "easy-dictionary-server/db"
	pointers "easy-dictionary-server/internalenv/utils"
	"time"
)

type UserEntity struct {
	ID         int                   `db:"id"`
	FirstName  string                `db:"first_name"`
	SecondName string                `db:"second_name"`
	Providers  *[]UserProviderEntity `db:"-"`
	CreatedAt  time.Time             `db:"created_at"`
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
	SecondName      string     `db:"second_name"`
	UserCreatedAt   time.Time  `db:"user_created_at"`
	ProviderID      *int       `db:"provider_id"`
	ProviderName    *string    `db:"provider_name"`
	Email           *string    `db:"email"`
	HashedPassword  *string    `db:"hashed_password"`
	ProviderCreated *time.Time `db:"provider_created_at"`
}

func GetAllUsers(db *database.Database, orderBy database.OrderByType) ([]UserEntity, error) {
	var rows []userWithProviderRow
	err := db.SQLDB.Select(&rows, GetAllUsersQuery(orderBy))
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
				ID:         row.UserID,
				FirstName:  row.FirstName,
				SecondName: row.SecondName,
				CreatedAt:  row.UserCreatedAt,
				Providers:  &[]UserProviderEntity{},
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
	err := db.SQLDB.Select(&rows, GetUserByIdQuery(), id)
	return mapUserWithProvidersToEntity(err, rows)
}

func GetUserByEmail(db *database.Database, email string) (*UserEntity, error) {
	var rows []userWithProviderRow
	err := db.SQLDB.Select(&rows, GetUserByEmailQuery(), email)
	return mapUserWithProvidersToEntity(err, rows)
}

func mapUserWithProvidersToEntity(err error, rows []userWithProviderRow) (*UserEntity, error) {
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	user := UserEntity{
		ID:         rows[0].UserID,
		FirstName:  rows[0].FirstName,
		SecondName: rows[0].SecondName,
		CreatedAt:  rows[0].UserCreatedAt,
		Providers:  &[]UserProviderEntity{},
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

func CreateUser(db *database.Database, user UserEntity) (int, error) {
	createdId := -1
	err := db.SQLDB.Select(&createdId, CreateUserQuery(), user.FirstName, user.SecondName,
		(*user.Providers)[0].ProviderName, (*user.Providers)[0].Email, (*user.Providers)[0].HashedPassword)
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
