package main

import "fmt"

// Database interface abstracts the database operations
type Database interface {
	Connect() string
	Query(query string) string
}

// MySQLDatabase is a concrete implementation of Database interface
type MySQLDatabase struct{}

func (db *MySQLDatabase) Connect() string {
	return "Connected to MySQL database."
}

func (db *MySQLDatabase) Query(query string) string {
	return fmt.Sprintf("MySQL Executed query: %s", query)
}

// PostgreSQLDatabase is another concrete implementation of Database interface
type PostgreSQLDatabase struct{}

func (db *PostgreSQLDatabase) Connect() string {
	return "Connected to PostgreSQL database."
}

func (db *PostgreSQLDatabase) Query(query string) string {
	return fmt.Sprintf("PostgreSQL Executed query: %s", query)
}

// DataManager uses the Database interface, protecting it from changes in database implementations
type DataManager struct {
	database Database
}

// NewDataManager initializes a new DataManager with a specified Database
func NewDataManager(db Database) *DataManager {
	return &DataManager{database: db}
}

// PerformOperations connects to the database and performs a sample query
func (dm *DataManager) PerformOperations() {
	fmt.Println(dm.database.Connect())
	fmt.Println(dm.database.Query("SELECT * FROM users;"))
}

func main() {
	mysqlDB := &MySQLDatabase{}
	postgresDB := &PostgreSQLDatabase{}

	// Using DataManager with MySQL database
	manager := NewDataManager(mysqlDB)
	manager.PerformOperations()

	// Switching to PostgreSQL database
	manager = NewDataManager(postgresDB)
	manager.PerformOperations()
}
