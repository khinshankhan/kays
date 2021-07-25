package usecases

import "time"

// Meta holds project metadata
type Meta struct {
	CommitHash string    `json:"commitHash"`
	BuildDate  string    `json:"buildDate"`
	SystemTime time.Time `json:"systemTime"`
}

// GetMeta collects all of the environment variables.
func GetMeta(BuildDate, CommitHash string) Meta {
	currentTimestamp := time.Now().UTC()

	meta := Meta{
		BuildDate:  BuildDate,
		CommitHash: CommitHash,
		SystemTime: currentTimestamp,
	}

	return meta
}
