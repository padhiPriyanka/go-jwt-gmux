package auth

import (
	"github.com/labstack/echo/v4"
	go_jwt_gmux "go-jwt-gmux"
	"go-jwt-gmux/pkg/api/auth/platform/pgsql"
	"gorm.io/gorm"
)

//New creates new iam service
func New(db *pg.DB, udb UserDB, j TokenGenerator, sec Securer, rbac RBAC) Auth {
	return Auth{
		db:   db,
		udb:  udb,
		tg:   j,
		sec:  sec,
		rbac: rbac,
	}
}

//Initialize initializes auth application service
func Initialize(db *pg.DB, j TokenGenerator, sec Securer, rbac RBAC) Auth {
	return New(db, pgsql.User{}, j, sec, rbac)
}

// Service represents auth service interface
type Service interface {
	Authenticate(echo.Context, string, string) (go_jwt_gmux.AuthToken, error)
	Refresh(echo.Context, string) (string, error)
	Me(echo.Context) (go_jwt_gmux.User, error)
}

// Auth represents auth application service
type Auth struct {
	db   *pg.DB
	udb  UserDB
	tg   TokenGenerator
	sec  Securer
	rbac RBAC
}

// UserDB represents user repository interface
type UserDB interface {
	View(gorm.DB, int) go_jwt_gmux.User
	//View(w http.ResponseWriter, r *http.Request) (go_jwt_gmux.User, error)
	FindByUsername(gorm.DB, string) (go_jwt_gmux.User, error)
	FindByToken(gorm.DB, string) (go_jwt_gmux.User, error)
	Update(gorm.DB, go_jwt_gmux.User) error
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(user go_jwt_gmux.User) (string, error)
}

// Securer represents security interface
type Securer interface {
	HashMatchesPassword(string, string) bool
	Token(string) string
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) go_jwt_gmux.AuthUser
}
