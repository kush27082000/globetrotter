package bussiness

import (
	"errors"
	"globa_trotter_game/dao"
	"globa_trotter_game/models"
	"regexp"
)

// Validate username: Only alphanumeric characters allowed
func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`) // Only letters and numbers
	return re.MatchString(username)
}

// Validate if the user exists before creating a new one
func RegisterUser(user *models.User) error {
	// Check if the username is empty
	// Check if username is valid
	isUserNameInValid, err := checkUserNameInvalid(user)
	if isUserNameInValid {
		return err
	}

	// Check if user already exists
	exists, err := dao.IsUserExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already taken")
	}

	// If not, create a new user in the DB
	return dao.RegisterUser(user)
}

func checkUserNameInvalid(user *models.User) (bool, error) {
	if user.Username == "" {
		return true, errors.New("username cannot be empty")
	}

	if !isValidUsername(user.Username) {
		return true, errors.New("username can only contain alphanumeric characters (letters and numbers)")
	}
	return false, nil
}

// LoginUser validates and logs in a user
func LoginUser(user *models.User) error {
	isUserNameInValid, err := checkUserNameInvalid(user)
	if isUserNameInValid {
		return err
	}
	// Check if the username exists
	exists, err := dao.IsUserExists(user.Username)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("invalid username")
	}
	return nil
}
