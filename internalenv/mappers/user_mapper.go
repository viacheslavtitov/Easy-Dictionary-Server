package mapper

import (
	dbUser "easy-dictionary-server/db/user"
	domainUser "easy-dictionary-server/domain/user"
)

func ToDomain(u *dbUser.UserEntity) *domainUser.User {
	return &domainUser.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		CreatedAt:  u.CreatedAt,
		Providers:  ToDomainProviders(u.Providers),
	}
}

func ToDomainProviders(u *[]dbUser.UserProviderEntity) *[]domainUser.UserProviders {
	if u == nil {
		return &[]domainUser.UserProviders{}
	}
	providers := make([]domainUser.UserProviders, len(*u))
	for i, item := range *u {
		providers[i] = *ToDomainProvider(&item)
	}
	return &providers
}

func ToDomainProvider(u *dbUser.UserProviderEntity) *domainUser.UserProviders {
	return &domainUser.UserProviders{
		ID:             u.ID,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		ProviderName:   u.ProviderName,
		CreatedAt:      u.CreatedAt,
	}
}

func FromDomain(u *domainUser.User) *dbUser.UserEntity {
	return &dbUser.UserEntity{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		CreatedAt:  u.CreatedAt,
		Providers:  FromDomainProviders(u.Providers, u.ID),
	}
}

func FromDomainProviders(u *[]domainUser.UserProviders, userId int) *[]dbUser.UserProviderEntity {
	if u == nil {
		return &[]dbUser.UserProviderEntity{}
	}
	providers := make([]dbUser.UserProviderEntity, len(*u))
	for i, item := range *u {
		providers[i] = *FromDomainProvider(&item, userId)
	}
	return &providers
}

func FromDomainProvider(u *domainUser.UserProviders, userId int) *dbUser.UserProviderEntity {
	return &dbUser.UserProviderEntity{
		ID:             u.ID,
		UserId:         userId,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		ProviderName:   u.ProviderName,
		CreatedAt:      u.CreatedAt,
	}
}
