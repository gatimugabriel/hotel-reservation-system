package constants

type Role string

const (
	ADMIN         Role = "ADMIN"
	PROPERTYOWNER Role = "PROPERTYOWNER"
	MANAGER       Role = "MANAGER"
	STAFF         Role = "STAFF"
	GUEST         Role = "GUEST"
)

type UserStatus string

const (
	ACTIVE   UserStatus = "ACTIVE"
	INACTIVE UserStatus = "INACTIVE"
	DELETED  UserStatus = "DELETED"
)