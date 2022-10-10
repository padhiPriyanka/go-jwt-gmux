package auth

import (
	"github.com/labstack/echo/v4"
	go_jwt_gmux "go-jwt-gmux"
	"net/http"
)

//Custom errors
var (
	ErrorInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Username or password does not exist")
)

// Authenticate tries to authenticate the user provided by username and password
func (a Auth) Authenticate(c echo.Context, user, pass string) (go_jwt_gmux.AuthToken, error) {
	u, err := a.udb.FindByUsername(a.db, user)
	if err != nil {
		return go_jwt_gmux.AuthToken{}, err
	}

	if !a.sec.HashMatchesPassword(u.Password, pass) {
		return go_jwt_gmux.AuthToken{}, ErrorInvalidCredentials
	}
	if !u.Active {
		return go_jwt_gmux.AuthToken{}, go_jwt_gmux.ErrUnauthorized
	}
	token, err := a.tg.GenerateToken(u)
	if err != nil {
		return go_jwt_gmux.AuthToken{}, go_jwt_gmux.ErrUnauthorized
	}
	u.updateLastLogin(a.sec.Token(token))
	if err := a.udb.Update(a.db, u); err != nil {
		return go_jwt_gmux.AuthToken{}, err
	}
	return go_jwt_gmux.AuthToken{Token: token, RefreshToken: u.Token}, nil
}

// Refresh refreshes jwt token and puts new claims inside
func (a Auth) Refresh(c echo.Context, refreshToken string) (string, error) {
	user, err := a.udb.FindByToken(a.db, refreshToken)
	if err != nil {
		return "", err
	}
	return a.tg.GenerateToken(user)
}

// Me returns info about currently logged user
func (a Auth) Me(c echo.Context) go_jwt_gmux.User {
	au := a.rbac.User(c)
	return a.udb.View(a.db, au.ID)
}
