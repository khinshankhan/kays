package config

// db defaults
const (
	// DefaultWebAddress is the default address the service look for the database
	DefaultDBHost = "localhost"

	// DefaultDBPort is the default port the service connect to the database
	DefaultDBPort = 5432

	// DefaultDBUser is the default user the service will try for the database
	DefaultDBUser = "postgres"

	// DefaultDBPassword is the default password the service will try for the database
	DefaultDBPassword = "postgres"

	// DefaultDBName is the default dbname the service connect to the database
	DefaultDBName = "caputo"
)

// web defaults
const (
	// DefaultWebAddress is the default address the service will bind to
	DefaultWebAddress = "0.0.0.0"

	// DefaultWebPort is the default port the service will listen on
	DefaultWebPort = 8080
)
