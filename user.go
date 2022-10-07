package go_jwt_gmux

import "time"

//user model

type User struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`

	//Mobile  string `json:"mobile,omitempty"`
	//Phone   string `json:"phone,omitempty"`
	//Address string `json:"address,omitempty"`

	Active bool `json:"active"`

	LastLogin          time.Time `json:"last_login,omitempty"`
	LastPasswordChange time.Time `json:"last_password_change,omitempty"`

	Token string `json:"-"`

	Role *Role `json:"role,omitempty"`

	RoleID     AccessRole `json:"-"`
	CompanyID  int        `json:"company_id"`
	LocationID int        `json:"location_id"`
}

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID         int
	CompanyID  int
	LocationID int
	Username   string
	Email      string
	Role       AccessRole
}
