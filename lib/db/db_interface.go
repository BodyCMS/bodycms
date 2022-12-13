package db

// Interface for the database
type IDnConnect interface {
	ConnectionString() string // Return the connection string
	Connect() error           // Connect to the database
	Close() error             // Close the connection to the database
}

// Interface for the database
type IDbAction interface {
	// Create a new record
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	FindMany(model interface{}, query interface{}, args ...interface{}) error
	FindOne(model interface{}, query interface{}, args ...interface{}) error
}
