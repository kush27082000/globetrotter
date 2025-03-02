package api

import (
	"fmt"
	"globa_trotter_game/bussiness"
	"globa_trotter_game/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call Business Layer to register user
	err := bussiness.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call Business Layer to validate login
	err := bussiness.LoginUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set Cookie on successful login
	c.SetCookie("username", user.Username, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// allowing any user to view clues
func GetRandomClue(c *gin.Context) {
	username := c.Query("username")
	fmt.Println("c.Cookieusername")
	fmt.Println(c.Cookie("username"))
	fmt.Println("username")
	fmt.Println(username)
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to play"})
		return
	}
	clueDetails, err := bussiness.GetRandomClueService(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": username, "clue": clueDetails})
}

// SubmitAnswer handles answer submission and score updating
func SubmitAnswer(c *gin.Context) {
	var request models.SubmitAnswerReq

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, destinationDetails, err := bussiness.SubmitAnswerService(request.Username, request.ClueID, request.Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response based on correct or incorrect answer
	if result == "correct" {
		c.JSON(http.StatusOK, gin.H{
			"result":    "correct",
			"city":      destinationDetails.City,
			"fun_facts": destinationDetails.FunFact,
			"trivia":    destinationDetails.Trivia,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "incorrect"})
	}
}

// GetScoreController retrieves the user's score.
func GetScoreController(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to view your score"})
		return
	}

	score, err := bussiness.GetScoreService(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": username, "score": score})
}

// ResetScoreController resets the user's score to 0.
func ResetScoreController(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to reset your score"})
		return
	}

	err := bussiness.ResetScoreService(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Score reset successfully", "username": username})
}

// CreateInviteController generates a shareable link.
func CreateInviteController(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to challenge a friend"})
		return
	}

	inviteLink, err := bussiness.GenerateInviteLink(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"invite_link": inviteLink})
}

// GetInviteController fetches inviter's score from the invite link.
func GetInviteController(c *gin.Context) {
	inviter := c.Param("username")

	score, err := bussiness.GetInviteDetails(inviter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"inviter": inviter,
		"score":   score,
		"message": inviter + " has challenged you! Try to beat their score.",
	})
}
