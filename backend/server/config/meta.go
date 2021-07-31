package config

// MetaConfig is the automatically filled meta information about the config
type MetaConfig struct {
	CommitHash string
	BuildDate  string

	// path to the file from which config was loaded from
	filePath string
}
