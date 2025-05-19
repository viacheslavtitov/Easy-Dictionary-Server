package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbUser "easy-dictionary-server/db/user"
	domain "easy-dictionary-server/domain/user"
	userMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

const (
	ErrUserNotFound = encodingError("User is not found")
)

type userRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) domain.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) (*domain.User, error) {
	zap.S().Debugf("Create user")
	userId, err := dbUser.CreateUser(ur.db, userMapper.FromDomain(user))
	if err != nil {
		return nil, err
	}
	user.ID = userId
	zap.S().Debugf("User created with %d id", userId)
	return user, nil
}

func (ur *userRepository) GetAllUsers(c context.Context) ([]*domain.User, error) {
	zap.S().Debugf("GetAllUsers")
	usersEntity, err := dbUser.GetAllUsers(ur.db, database.OrderByASC)
	users := []*domain.User{}
	for _, item := range usersEntity {
		users = append(users, userMapper.ToDomain(&item))
	}
	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	zap.S().Debugf("GetByEmail %s", email)
	userEntity, err := dbUser.GetUserByEmail(ur.db, email)
	if err != nil {
		return nil, err
	}
	return userMapper.ToDomain(userEntity), nil
}

func (ur *userRepository) GetByID(c context.Context, id int) (*domain.User, error) {
	zap.S().Debugf("GetByID %d", id)
	userEntity, err := dbUser.GetUserById(ur.db, id)
	if err != nil {
		return nil, err
	}
	return userMapper.ToDomain(userEntity), nil
}

func (ur *userRepository) UpdateUser(c context.Context, user *domain.User) (*domain.User, error) {
	zap.S().Debugf("UpdateUser %s %s", user.FirstName, user.SecondName)
	userEntity := userMapper.FromDomain(user)
	err := dbUser.UpdateUser(ur.db, *userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(c context.Context, id int) error {
	zap.S().Debugf("DeleteUser %d", id)
	err := dbUser.DeleteUserById(ur.db, id)
	if err != nil {
		return err
	}
	return nil
}
