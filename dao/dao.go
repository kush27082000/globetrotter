package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"globa_trotter_game/models"
	"globa_trotter_game/utils/database"
)

// Check if the user exists in the database
func IsUserExists(username string) (bool, error) {
	db := database.GetDB()
	var count int

	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // User does not exist
		}
		return false, err
	}

	return count > 0, nil
}

// Insert a new user into the database
func RegisterUser(user *models.User) error {
	db := database.GetDB()

	// Prepare SQL statement
	query := `INSERT INTO users (username) VALUES (?)`
	_, err := db.Exec(query, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func GetRandomDestination() (models.ClueDetails, error) {
	db := database.GetDB() // Get *sql.DB connection

	var destination models.ClueDetails

	query := "SELECT id, clues FROM destinations ORDER BY RAND() LIMIT 1"

	row := db.QueryRow(query)

	err := row.Scan(&destination.ID, &destination.Clues)
	if err != nil {
		if err == sql.ErrNoRows {
			return destination, fmt.Errorf("no destinations found")
		}
		return destination, err
	}

	return destination, nil
}

// GetDestinationDetails fetches city, fun facts, and trivia for the given clue ID
func GetDestinationDetails(clueID uint) (models.DestinationDetails, error) {
	db := database.GetDB()
	var details models.DestinationDetails
	query := "SELECT city, fun_facts, trivia FROM destinations WHERE id = ?"
	err := db.QueryRow(query, clueID).Scan(&details.City, &details.FunFact, &details.Trivia)
	if err != nil {
		if err == sql.ErrNoRows {
			return details, errors.New("clue not found")
		}
		return details, err
	}
	return details, nil
}

// UpdateUserScore increments the user's score
func UpdateUserScore(username string) error {
	db := database.GetDB()
	query := "UPDATE users SET score = score + 1 WHERE username = ?"
	_, err := db.Exec(query, username)
	return err
}

// GetUserScore retrieves the score for a given username.
func GetUserScore(username string) (int, error) {
	db := database.GetDB()
	var score int
	query := "SELECT score FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&score)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}
	return score, nil
}

// ResetUserScore sets the user's score to 0.
func ResetUserScore(username string) error {
	db := database.GetDB()
	query := "UPDATE users SET score = 0 WHERE username = ?"
	_, err := db.Exec(query, username)
	return err
}
