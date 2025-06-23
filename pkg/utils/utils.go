package utils

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)



func GenerateOtp(length int) string {
	if length < 1 {
		length = 4
	}



	rand.Seed(time.Now().UnixNano())

	otp := ""

	for i := 0; i < length; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp

}

type UserModel string

const (
	User  UserModel = "user"
	Admin UserModel = "admin"
)


func ToNullString(s *string) sql.NullString {
	if s != nil && *s != "" {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}
func ToNullBool(b *bool) sql.NullBool {
	if b != nil {
		return sql.NullBool{
			Bool:  *b,
			Valid: true,
		}
	}
	return sql.NullBool{
		Valid: false,
	}
}