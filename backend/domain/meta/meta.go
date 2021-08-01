package meta

import "time"

// Meta holds project metadata
type Meta struct {
	CommitHash string    `json:"commitHash"`
	BuildDate  string    `json:"buildDate"`
	SystemTime time.Time `json:"systemTime"`
}
