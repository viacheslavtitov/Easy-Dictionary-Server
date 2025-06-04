package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbUser "easy-dictionary-server/db/user"
	domain "easy-dictionary-server/domain/user"
	userMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type userRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) domain.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) (*domain.User, error) {
	zap.S().Debugf("Create user")
	uuid, err := dbUser.CreateUser(ur.db, userMapper.FromUserDomain(user, nil))
	if err != nil {
		return nil, err
	}
	user.UUID = *uuid
	zap.S().Debugf("User created with %d uuid", uuid)
	return user, nil
}

func (ur *userRepository) GetAllUsers(c context.Context) ([]*domain.User, error) {
	zap.S().Debugf("GetAllUsers")
	usersEntity, err := dbUser.GetAllUsers(ur.db, database.OrderByASC)
	users := []*domain.User{}
	for _, item := range usersEntity {
		users = append(users, userMapper.ToUserDomain(&item))
	}
	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (*domain.User, *int, error) {
	zap.S().Debugf("GetByEmail %s", email)
	userEntity, err := dbUser.GetUserByEmail(ur.db, email)
	// zap.S().Debugf("User UUID %s", userEntity.UUID)
	if err != nil {
		return nil, nil, err
	}
	return userMapper.ToUserDomain(userEntity), &userEntity.ID, nil
}

func (ur *userRepository) GetByID(c context.Context, id int) (*domain.User, error) {
	zap.S().Debugf("GetByID %d", id)
	userEntity, err := dbUser.GetUserById(ur.db, id)
	if err != nil {
		return nil, err
	}
	return userMapper.ToUserDomain(userEntity), nil
}

func (ur *userRepository) GetByUUID(c context.Context, uuid string) (*domain.User, error) {
	zap.S().Debugf("GetByUUID %s", uuid)
	userEntity, err := dbUser.GetUserByUUID(ur.db, uuid)
	if err != nil {
		return nil, err
	}
	return userMapper.ToUserDomain(userEntity), nil
}

func (ur *userRepository) UpdateUser(c context.Context, user *domain.User, userId int) (*domain.User, error) {
	zap.S().Debugf("UpdateUser %s %s", user.FirstName, user.LastName)
	userEntity := userMapper.FromUserDomain(user, &userId)
	err := dbUser.UpdateUser(ur.db, userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteUser %d", id)
	rowsDeleted, errQuery := dbUser.DeleteUserById(ur.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
