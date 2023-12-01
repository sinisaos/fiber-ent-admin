package main

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/sinisaos/fiber-ent-admin/pkg/database"
	"github.com/sinisaos/fiber-ent-admin/pkg/utils"

	"github.com/manifoldco/promptui"
)

func main() {
	// Create ent client
	client := database.DbConnection()
	var resultSuperUserBoolean bool

	// Get the username
	promptUserName := promptui.Prompt{
		Label: "Username",
	}

	resultUserName, err := promptUserName.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Validate the email
	validateEmail := func(email string) error {
		err, _ := mail.ParseAddress(email)
		if err == nil {
			return errors.New("invalid email")
		}
		return nil
	}

	// Get the email
	promptEmail := promptui.Prompt{
		Label:    "Email",
		Validate: validateEmail,
	}

	resultEmail, err := promptEmail.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Validate the password
	validatePasword := func(input string) error {
		if len(input) < 6 {
			return errors.New("password must have more than 6 characters")
		}
		return nil
	}

	// Get the password
	promptPassword := promptui.Prompt{
		Label:    "Password",
		Validate: validatePasword,
		Mask:     '*',
	}

	resultPassword, err := promptPassword.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Validate matched password
	validateConfirmPasword := func(input string) error {
		if input != resultPassword {
			return errors.New("password does not match")
		}
		return nil
	}

	// Get the confirmed password
	promptConfirmPassword := promptui.Prompt{
		Label:    "Confirm password",
		Validate: validateConfirmPasword,
		Mask:     '*',
	}

	resultConfirmPassword, err := promptConfirmPassword.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Get the superuser
	promptSuperUser := promptui.Prompt{
		Label: "Superuser (y/n)",
	}

	resultSuperUser, err := promptSuperUser.Run()

	if resultSuperUser == "y" {
		resultSuperUserBoolean = true
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	hashedPassword, err := utils.HashPassword(resultConfirmPassword)
	if err != nil {
		return
	}

	// Create user
	user, err := client.User.Create().
		SetEmail(resultEmail).
		SetUsername(resultUserName).
		SetPassword(hashedPassword).
		SetSuperuser(resultSuperUserBoolean).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return
	}

	fmt.Printf("User %v successfully created\n", user.ID)
}
