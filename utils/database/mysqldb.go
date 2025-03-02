package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MSSQLConfig struct {
	Server                string        `json:"server"`
	Port                  int           `json:"port"`
	Name                  string        `json:"name"`
	Username              string        `json:"-"`
	Password              string        `json:"-"`
	Table                 string        `json:"-"`
	MaxOpenConnections    int           `json:"maxOpenConnections"`
	MaxIdleConnections    int           `json:"maxIdleConnections"`
	ConnectionMaxLifetime time.Duration `json:"connectionMaxLifetime"`
	ConnectionMaxIdleTime time.Duration `json:"connectionMaxIdleTime"`
}

var (
	mysqldb *sql.DB
	once    sync.Once
)

// Initialize the database connection
func InitDBWithConfig(mysqlConfig MSSQLConfig) error {
	once.Do(func() {
		// Form the data source name (DSN)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Server, mysqlConfig.Port, mysqlConfig.Name)
		fmt.Println("dsn is ", dsn)

		// Open a connection to the database
		var err error
		mysqldb, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		// Set database connection pool settings
		mysqldb.SetMaxOpenConns(mysqlConfig.MaxOpenConnections)
		mysqldb.SetMaxIdleConns(mysqlConfig.MaxIdleConnections)
		mysqldb.SetConnMaxLifetime(mysqlConfig.ConnectionMaxLifetime)
		mysqldb.SetConnMaxIdleTime(mysqlConfig.ConnectionMaxIdleTime)

		// Ping the database to verify connection
		err = mysqldb.Ping()
		if err != nil {
			mysqldb.Close() // Close the database if ping fails
			log.Fatalf("Error connecting to the database: %v", err)
		}

		fmt.Println("Database successfully initialized!")
	})
	return nil
}

// GetDB returns the singleton database connection
func GetDB() *sql.DB {
	if mysqldb == nil {
		log.Fatal("Database not initialized! Call InitDBWithConfig first.")
	}
	return mysqldb
}
