package dbaxs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

/*
 * Connect to the database
 * Returns true if successfully connected
 */
func ConnDB(username string, password string, ipAddress string, port string, dbName string) bool {

	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, ipAddress, port, dbName)
	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
		return false
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return false
	}
	fmt.Println("Successfully reached DB!")
	DB.SetMaxOpenConns(10000)
	DB.SetConnMaxLifetime(time.Second * 60)
	fmt.Println("max open connections:", DB.Stats().MaxOpenConnections)
	return true
}
