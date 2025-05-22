package mapper

import (
	dbUser "easy-dictionary-server/db/user"
	domainUser "easy-dictionary-server/domain/user"
)

func ToUserDomain(u *dbUser.UserEntity) *domainUser.User {
	return &domainUser.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		Role:      u.Role,
		Providers: ToUserDomainProviders(u.Providers),
	}
}

func ToUserDomainProviders(u *[]dbUser.UserProviderEntity) *[]domainUser.UserProviders {
	if u == nil {
		return &[]domainUser.UserProviders{}
	}
	providers := make([]domainUser.UserProviders, len(*u))
	for i, item := range *u {
		providers[i] = *ToUserDomainProvider(&item)
	}
	return &providers
}

func ToUserDomainProvider(u *dbUser.UserProviderEntity) *domainUser.UserProviders {
	return &domainUser.UserProviders{
		ID:             u.ID,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		ProviderName:   u.ProviderName,
		CreatedAt:      u.CreatedAt,
	}
}

func FromUserDomain(u *domainUser.User) *dbUser.UserEntity {
	return &dbUser.UserEntity{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		Providers: FromUserDomainProviders(u.Providers, u.ID),
		Role:      u.Role,
	}
}

func FromUserDomainProviders(u *[]domainUser.UserProviders, userId int) *[]dbUser.UserProviderEntity {
	if u == nil {
		return &[]dbUser.UserProviderEntity{}
	}
	providers := make([]dbUser.UserProviderEntity, len(*u))
	for i, item := range *u {
		providers[i] = *FromUserDomainProvider(&item, userId)
	}
	return &providers
}

func FromUserDomainProvider(u *domainUser.UserProviders, userId int) *dbUser.UserProviderEntity {
	return &dbUser.UserProviderEntity{
		ID:             u.ID,
		UserId:         userId,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		ProviderName:   u.ProviderName,
		CreatedAt:      u.CreatedAt,
	}
}
