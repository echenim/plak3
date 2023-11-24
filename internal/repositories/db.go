package repositories

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib" // PostgreSQL driver
)

func InitDatabase(driver, dataSourceName string) (*sql.DB, error) {
	// Open the database connection
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// var DB *sql.DB

// func InitDatabase(driver, dataSourceName string) error {
// 	var err error
// 	DB, err = sql.Open(driver, dataSourceName)
// 	if err != nil {
// 		return err
// 	}
// 	return DB.Ping()
// }
