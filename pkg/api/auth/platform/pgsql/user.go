package pgsql

import (
	go_jwt_gmux "go-jwt-gmux"
	"gorm.io/gorm"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db gorm.DB, id int) go_jwt_gmux.User {
	var user go_jwt_gmux.User
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."id" = ? and deleted_at is null)`
	db.Find(&user, sql, id)
	return user
}
