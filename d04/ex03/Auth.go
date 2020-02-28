package Ex03

import (
	"crypto/subtle"
	Ex00 "d04/ex00"
	Ex02 "d04/ex02"
	"fmt"
)

func Auth(login, password string) (valid bool) {
	var current *Ex00.User
	for _, user := range Ex00.Store {
		if subtle.ConstantTimeCompare([]byte(login), []byte(user.Login)) == 1 {
			current = user
			break
		}
	}
	if current == nil {
		return false
	}
	fmt.Println("Comparing passwords", password, current.Password)
	valid = Ex02.ComparePasswords(current.Password, []byte(password))
	return valid
}
