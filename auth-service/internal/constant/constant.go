package constant

import (
	userstatus "github.com/truongle2004/service-context/core"
)

// API versions
const (
	V1 = "/api/v1"
	V2 = "/api/v2"
)

// Default values and formats
const (
	DefaultUserRole = "ROLE_USER"
	DateFormat      = "2006-01-02"
	SaltLimit       = 16
)

// Authentication-related constants
const (
	TokenType           = "Bearer"
	HeaderAuthorization = "Authorization"
	XAuthUserID         = "X-Auth-User-Id"
	XAuthUserStatus     = "X-Auth-User-Status"
)

// Messages
const (
	MessageInvalidCredential = "Invalid email and password"
)

// UserStatus represents the status of a user
type UserStatus string

const (
	UserStatusActive  UserStatus = UserStatus(userstatus.ACTIVE)
	UserStatusReject  UserStatus = UserStatus(userstatus.REJECT)
	UserStatusUnBlock UserStatus = UserStatus(userstatus.UNLOCK)
	UserStatusDeleted UserStatus = UserStatus(userstatus.DELETE)
	UserStatusLocked  UserStatus = UserStatus(userstatus.LOCKED)
	UserStatusPending UserStatus = UserStatus(userstatus.PENDING)
)
