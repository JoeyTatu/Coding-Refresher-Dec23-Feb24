package storage

import (
	"fmt"

	"github.com/jackc/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	// Create the database if it doesn't exist
	err := createDatabaseIfNotExists(config)
	if err != nil {
		return nil, err
	}

	// Now, establish a connection to the database
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}

func createDatabaseIfNotExists(config *Config) error {
    // Connect to PostgreSQL server without specifying a specific database
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s sslmode=%s",
        config.Host, config.Port, config.User, config.Password, config.SSLMode,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    // Close the underlying sql.DB connection when done
    sqlDB, err := db.DB()
    if err != nil {
        return err
    }
    defer sqlDB.Close()

    // Attempt to create the database
    createDBQuery := fmt.Sprintf("CREATE DATABASE %s", config.DBName)
    if err := db.Exec(createDBQuery).Error; err != nil {
        // Handle the case when the database already exists
        if isDatabaseExistsError(err) {
            fmt.Printf("Database %s already exists\n", config.DBName)
            return nil
        }
        return err
    }

    fmt.Printf("Database %s created successfully\n", config.DBName)
    return nil
}

func isDatabaseExistsError(err error) bool {
    pgErr, ok := err.(*pgconn.PgError)
    if !ok {
        return false
    }
    // SQLSTATE 42P04 corresponds to "duplicate_database" in PostgreSQL
    return pgErr.SQLState() == "42P04"
}
