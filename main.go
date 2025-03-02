package main

import (
	"context"
	"fmt"
	"globa_trotter_game/api"
	"globa_trotter_game/constants"
	"globa_trotter_game/utils/configs"
	"globa_trotter_game/utils/database"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	initConfigs()
	r := gin.Default()
	// Enable CORS
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	config := cors.Config{
		AllowOrigins:     []string{"*"}, // Change to your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
		MaxAge:           12 * time.Hour, // Reduce preflight requests
	}
	r.Use(cors.New(config))

	r.POST("/register", api.RegisterUser)
	r.POST("/login", api.LoginUser)
	r.GET("/clue", api.GetRandomClue)
	r.POST("/submit", api.SubmitAnswer)
	r.GET("/score", api.GetScoreController)
	r.POST("/reset", api.ResetScoreController)
	r.POST("/invite", api.CreateInviteController)       // Generate link
	r.GET("/invite/:username", api.GetInviteController) // Fetch inviter's score
	r.Run(":8080")
}

func initConfigs() {
	configs.Init(constants.BaseConfigPathUATPtrms)
	var ctx context.Context
	initMySqlDB(ctx)
	log.Print("initialized configs")
}

func initMySqlDB(ctx context.Context) {
	// init database
	mysqldbConfig, err := configs.Get(constants.MySQL_DB_Config)
	if err != nil {
		log.Print("error getting mssql database config ")
	}

	fmt.Print("constants.MySQL_DB_Config   ")
	fmt.Println(constants.MySQL_DB_Config)

	err = database.InitDBWithConfig(database.MSSQLConfig{
		Server:                mysqldbConfig.GetString("mysql_db.server"),
		Port:                  mysqldbConfig.GetInt("mysql_db.port"),
		Name:                  mysqldbConfig.GetString("mysql_db.name"),
		Username:              mysqldbConfig.GetString("mysql_db.username"),
		Password:              mysqldbConfig.GetString("mysql_db.password"),
		Table:                 mysqldbConfig.GetString("mysql_db.table"),
		MaxOpenConnections:    mysqldbConfig.GetInt("mysql_db.maxOpenConnections"),
		MaxIdleConnections:    mysqldbConfig.GetInt("mysql_db.maxIdleConnections"),
		ConnectionMaxLifetime: mysqldbConfig.GetDuration("mysql_db.connectionMaxLifetimeInSeconds") * time.Second,
		ConnectionMaxIdleTime: mysqldbConfig.GetDuration("mysql_db.connectionMaxIdleTimeInSeconds") * time.Second,
	})

	if err != nil {
		log.Print("unable to initialize database")
	}

	log.Print("initialized database")
}
