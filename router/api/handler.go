package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/sanjivyash/AuthAPI/models/token"
	"github.com/sanjivyash/AuthAPI/models/user"
)

// signup endpoint
func createUser(c *fiber.Ctx) error {
	if !c.Is("json") {
		return c.Status(400).JSON(map[string]string {"error": "Only JSON is accepted..."})
	}

	var user user.User

	// username and password field entered
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(user)
		return c.Status(500).JSON(map[string]string{"error": "Please input valid fields username and password"})
	}

	// user signup
	if err := user.Save(); err != nil {
		return c.Status(400).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(201).JSON(user)
}

// login endpoint
func loginUser(c *fiber.Ctx) error {
	if !c.Is("json") {
		return c.Status(400).JSON(map[string]string {"error": "Only JSON is accepted..."})
	}

	var user user.User

	// username and password field entered
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(user)
		return c.Status(500).JSON(map[string]string{"error": "Please input valid fields username and password"})
	}

	// login authentication
	if err := user.Login(); err != nil {
		return c.Status(401).JSON(map[string]string{"error": err.Error()})
	}

	// user authenticated - generate token
	var token token.Token
	token.Generate()

	// return c.SendString(token.Message)
	return c.JSON(map[string]string{"username": user.Username, "token": token.Message})
}

// delete user endpoint
func deleteUser(c *fiber.Ctx) error {
	if !c.Is("json") {
		return c.Status(400).JSON(map[string]string {"error": "Only JSON is accepted..."})
	}

	var user user.User

	// username and password field entered
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(user)
		return c.Status(500).JSON(map[string]string{"error": "Please input valid fields username and password"})
	}

	// login authentication
	if err := user.Login(); err != nil {
		return c.Status(401).JSON(map[string]string{"error": err.Error()})
	}

	// user authenticated - delete user
	user.Delete()
	return c.JSON(user)
}

// get info
func getInfo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "secret revealed haha"})
}