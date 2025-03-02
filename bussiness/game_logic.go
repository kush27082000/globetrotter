package bussiness

import (
	"errors"
	"fmt"
	"globa_trotter_game/dao"
	"globa_trotter_game/models"
)

// GetRandomClueService fetches a random clue after validating the username
func GetRandomClueService(username string) (models.ClueDetails, error) {
	user := &models.User{Username: username}

	// Check if username is invalid
	isUserNameInValid, err := checkUserNameInvalid(user)
	if isUserNameInValid {
		return models.ClueDetails{}, err
	}

	// Check if the username exists in the database
	exists, err := dao.IsUserExists(user.Username)
	if err != nil {
		return models.ClueDetails{}, err
	}
	if !exists {
		return models.ClueDetails{}, errors.New("invalid username")
	}

	// Fetch a random clue from the database
	return dao.GetRandomDestination()
}

// SubmitAnswerService validates the answer and updates the score if correct
func SubmitAnswerService(username string, clueID uint, answer string) (string, models.DestinationDetails, error) {
	// Check if the username exists
	exists, err := dao.IsUserExists(username)
	if err != nil {
		return "", models.DestinationDetails{}, err
	}
	if !exists {
		return "", models.DestinationDetails{}, errors.New("invalid username")
	}

	// Get correct city, fun facts, and trivia from the database
	destinationDetails, err := dao.GetDestinationDetails(clueID)
	if err != nil {
		return "", models.DestinationDetails{}, err
	}

	// Compare answer with the correct city
	if answer == destinationDetails.City {
		// Increase user score
		err := dao.UpdateUserScore(username)
		if err != nil {
			return "", models.DestinationDetails{}, err
		}
		return "correct", destinationDetails, nil
	}
	return "incorrect", models.DestinationDetails{}, nil
}

// GetScoreService fetches the user's score.
func GetScoreService(username string) (int, error) {
	// Check if the user exists
	exists, err := dao.IsUserExists(username)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errors.New("user not found")
	}

	// Fetch score
	return dao.GetUserScore(username)
}

// ResetScoreService resets the user's score to 0.
func ResetScoreService(username string) error {
	// Check if the user exists
	exists, err := dao.IsUserExists(username)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user not found")
	}

	// Reset the score
	return dao.ResetUserScore(username)
}

// GenerateInviteLink creates a unique shareable invite link.
func GenerateInviteLink(inviter string) (string, error) {
	// Check if the inviter exists
	exists, err := dao.IsUserExists(inviter)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("user not found")
	}

	// Generate dynamic invite link
	inviteLink := fmt.Sprintf("http://localhost:8080/invite/%s", inviter)
	return inviteLink, nil
}

// GetInviteDetails fetches the inviter's score.
func GetInviteDetails(inviter string) (int, error) {
	score, err := dao.GetUserScore(inviter)
	if err != nil {
		return 0, err
	}
	return score, nil
}
