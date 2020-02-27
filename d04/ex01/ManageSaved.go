package Ex01

import (
	"bytes"
	Ex00 "d04/ex00"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

func transformUsersArrToMap(store map[string]*Ex00.User, users [][]string) {
	for _, elem := range users {
		login := elem[0]
		_, ok := store[login]
		if !ok {
			store[login] = &Ex00.User{Login: login, Password: elem[1]}
		}
	}
}

func transformUsersMapToArr(users map[string]*Ex00.User) (arr [][]string) {
	arr = make([][]string, 0)
	for _, user := range users {
		elem := make([]string, 2)
		elem[0] = user.Login
		elem[1] = user.Password
		arr = append(arr, elem)
	}
	return
}

func writeUsersToFile(file, str string) (err error) {
	err = ioutil.WriteFile(file, []byte(str), 0666)
	if err != nil {
		fmt.Println("Failed to open ../private/passwd to save passwords:", err)
	}
	return
}

func deserializeArr(by []byte, arr *[][]string) (err error) {
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(arr)
	return
}

func decode(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func SaveUsers(users map[string]*Ex00.User) (err error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	arr := transformUsersMapToArr(users)
	if err = encoder.Encode(arr); err != nil {
		fmt.Println("Failed to Serialize user data to save:", err)
		return
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	return writeUsersToFile(FILE, str)
}

func LoadUsers() (err error) {
	var by []byte
	var arr [][]string
	if data, err := ioutil.ReadFile(FILE); err != nil {
		fmt.Println("Error Loading Users:", err)
	} else if by, err = decode(data); err != nil {
		fmt.Println("Error Loading Users:", err)
	} else if err = deserializeArr(by, &arr); err != nil {
		fmt.Println("Failed to deserialize passwords file:", err)
	} else {
		transformUsersArrToMap(Ex00.Store, arr)
	}
	return
}
