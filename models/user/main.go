package user

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/sanjivyash/AuthAPI/config"
	"github.com/sanjivyash/AuthAPI/database"
)

var path string = config.Config("BASE_DIR") + "/storage/users.json"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// salt the password
func (user *User) Salt() error {
	if user.Username == "" || user.Password == "" {
		return errors.New("Please enter username and password fields")
	}
	h := sha256.New()
	h.Write([]byte(user.Password))

	user.Password = fmt.Sprintf("%x", h.Sum(nil))
	return nil
}

// save user in database
func (user User) Save() error {

	data := database.ReadFile(path)
	fmt.Println("Successfully opened users file")

	var users []User
	json.Unmarshal(data, &users)

	for i := 0; i < len(users); i++ {
		if users[i].Username == user.Username {
			fmt.Println("Username already in use")
			return errors.New("Username already in use")
		}
	}

	users = append(users, user)
	updata, err := json.Marshal(users)
	fmt.Println([]byte(user.Password))

	if err != nil {
		log.Fatal("[ERROR] Error in converting to JSON\n" + err.Error())
	}

	return database.WriteFile(path, updata)
}

// login existing user
func (user User) Login() error {
	data := database.ReadFile(path)

	fmt.Println("Successfully opened users file")
	var users []User

	json.Unmarshal(data, &users)

	for i := 0; i < len(users); i++ {
		if users[i] == user {
			fmt.Println("Welcome User: " + user.Username)
			return nil
		} else {
			fmt.Println([]byte(user.Password))
			fmt.Println(user.Password)
			fmt.Println([]byte(users[i].Password))
			fmt.Println(users[i].Password)
		}
	}

	fmt.Println("Invalid Credentials")
	return errors.New("Invalid Credentials")
}

// delete user from database
func (user User) Delete() error {
	data := database.ReadFile(path)

	fmt.Println("Successfully opened users file")
	var users []User

	json.Unmarshal(data, &users)

	for i := 0; i < len(users); i++ {
		if users[i] == user {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	updata, err := json.Marshal(users)

	if err != nil {
		log.Fatal("[ERROR] Error in converting to JSON\n" + err.Error())
	}

	return database.WriteFile(path, updata)
}
