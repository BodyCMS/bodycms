package db

import "fmt"

type DbType string

const (
	DbTypeMySQL    DbType = "mysql"
	DbTypePostgres DbType = "postgres"
	DbTypeSQLite   DbType = "sqlite"
)

// Options for the database connection
// This options are used to connect to the database
//
// Now supported: MySQL, PostgreSQL, SQLite (default)
type DbOptions struct {
	Host             string `json:"host"`              // Hostname or IP address
	Port             string `json:"port"`              // Port number
	Username         string `json:"username"`          // Username
	Password         string `json:"password"`          // Password
	Database         string `json:"database"`          // Database name
	SllMode          string `json:"sslmode"`           // SSL mode
	ConnectionString string `json:"connection_string"` // Connection string, if this is set, the other options will be ignored
	DbType           DbType `json:"db_type"`           // Database type
}

// Return the connection string based on the options
func (o *DbOptions) GetConnectionString() string {
	if o.ConnectionString != "" {
		return o.ConnectionString
	}

	switch o.DbType {
	case DbTypeMySQL:
		return o.getMySQLConnectionString()
	case DbTypePostgres:
		return o.getPostgresConnectionString()
	case DbTypeSQLite:
		return o.getSQLiteConnectionString()
	}

	return ""
}

func (o *DbOptions) getMySQLConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", o.Username, o.Password, o.Host, o.Port, o.Database)
}

func (o *DbOptions) getPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", o.Host, o.Port, o.Username, o.Password, o.Database, o.SllMode)
}

func (o *DbOptions) getSQLiteConnectionString() string {
	return fmt.Sprintf("%s.db", o.Database)
}

// Get the default database connection, MySql. "bodycms.db"
func (o *DbOptions) GetDefaultDb() string {
	return "bodycms.db"
}
