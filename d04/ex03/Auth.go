package Ex03

import (
	"crypto/subtle"
	Ex00 "d04/ex00"
	Ex02 "d04/ex02"
	"fmt"
)

func Auth(login, password, sessionUUID string) (valid bool) {
	current := Ex00.Store[sessionUUID]
	fmt.Println("comparing:", login, password, current)
	if current == nil || subtle.ConstantTimeCompare([]byte(login), []byte(current.Login)) != 1 {
		return false
	}
	valid = Ex02.ComparePasswords(password, []byte(current.Password))
	return valid
}
