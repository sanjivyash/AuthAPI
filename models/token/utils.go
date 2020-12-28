package token 

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/sanjivyash/AuthAPI/database"
)

// utility function to generate token message
func generateToken() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	chars := make([]byte, tokenLength)
  for i := range chars {
    chars[i] = charset[rand.Intn(len(charset))]
  }

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