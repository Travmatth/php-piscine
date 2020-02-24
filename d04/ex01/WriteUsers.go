package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

func transformUsersArrToMap(store map[string]*User, users [][]string) {
	for _, elem := range users {
		login := elem[0]
		_, ok := store[login]
		if !ok {
			store[login] = &User{Login: login, Password: elem[1]}
		}
	}
}

func transformUsersMapToArr(users map[string]*User) (arr [][]string) {
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

func SaveUsers(users map[string]*User) (err error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	arr := transformUsersMapToArr(users)
	if err = encoder.Encode(arr); err != nil {
		fmt.Println("Failed to Serialize user data to save:", err)
		return
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	return writeUsersToFile(file, str)
}

func LoadUsers() (err error) {
	var by []byte
	var arr [][]string
	if data, err := ioutil.ReadFile(file); err != nil {
		fmt.Println("Error Loading Users:", err)
	} else if by, err = decode(data); err != nil {
		fmt.Println("Error Loading Users:", err)
	} else if err = deserializeArr(by, &arr); err != nil {
		fmt.Println("Failed to deserialize passwords file:", err)
	} else {
		transformUsersArrToMap(store, arr)
	}
	return
}
