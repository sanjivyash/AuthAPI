package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/sanjivyash/AuthAPI/config"
	"github.com/sanjivyash/AuthAPI/database"
)

var path string = config.Config("BASE_DIR") + "/storage/tokens.json"
var tokenLength, _ = strconv.ParseInt(config.Config("TOKEN_LENGTH"), 10, 64)
var timeLimit, _ = strconv.ParseInt(config.Config("TIME_LIMIT"), 10, 64)

type Token struct {
	CreatedAt int64  `json:"createdAt"`
	Message   string `json:"token"`
}

// generate token for logging user
func (token *Token) Generate() {
	tokens := updatedTokens()
	token.CreatedAt = time.Now().Unix()
	token.Message = generateToken()

	tokens = append(tokens, *token)
	data, err := json.Marshal(tokens)

	if err != nil {
		log.Fatal("[ERROR] Error in converting to JSON\n" + err.Error())
	}

	if database.WriteFile(path, data) != nil {
		log.Fatal("[ERROR] Unable to write in tokens file")
	}
}

// authenticate info request
func (token Token) Authenticate() error {
	tokens := updatedTokens()
	fmt.Println(tokens)

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Message == token.Message {
			fmt.Println("Valid Token: " + token.Message)
			return nil
		}
	}

	fmt.Println("Invalid Token")
	return errors.New("Invalid Token")
}

// utility function to generate token message
func generateToken() string {
	chars := make([]byte, tokenLength)
	rand.Read(chars)

	msg := string(chars)
	fmt.Println(msg)
	return msg
}

// utility function to remove expired tokens
func updatedTokens() []Token {
	data := database.ReadFile(path)
	fmt.Println("Successfully opened tokens file")

	var tokens, validTokens []Token
	json.Unmarshal(data, &tokens)

	for i := 0; i < len(tokens); i++ {
		if time.Now().Unix()-tokens[i].CreatedAt < timeLimit {
			validTokens = append(validTokens, tokens[i])
		}
	}

	data, err := json.Marshal(validTokens)

	if err != nil {
		log.Fatal("[ERROR] Error in converting to JSON\n" + err.Error())
	}

	if err = database.WriteFile(path, data); err != nil {
		log.Fatal("[ERROR] Unable to write in tokens file\n" + err.Error())
	}

	return validTokens
}
