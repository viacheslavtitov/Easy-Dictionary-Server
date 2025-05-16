package mapper

import (
	dbUser "easy-dictionary-server/db/user"
	domainUser "easy-dictionary-server/domain/user"
)

func ToDomain(u *dbUser.UserEntity) *domainUser.User {
	return &domainUser.User{
		ID:         u.ID,
		Email:      u.Email,
		ProviderId: u.ProviderID,
		UID:        u.UID,
	}
}

func FromDomain(u *domainUser.User) *dbUser.UserEntity {
	return &dbUser.UserEntity{
		ID:         u.ID,
		Email:      u.Email,
		ProviderID: u.ProviderId,
		UID:        u.UID,
	}
}
