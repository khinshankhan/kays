package usecases

import (
	"time"

	"github.com/kkhan01/caputo/backend/config"
	metadomain "github.com/kkhan01/caputo/backend/domain/meta"
)

type (
	// Usecases declares available services
	Usecases interface {
		GetMeta() metadomain.Meta
	}

	// usecases declares the dependencies for the service
	usecases struct {
		metaConfig *config.MetaConfig
	}
)

// NewUsecases returns Usecases
func NewUsecases(metaConfig *config.MetaConfig) usecases {
	return usecases{
		metaConfig: metaConfig,
	}
}

// GetMeta returns all of the environment variables
func (usecases usecases) GetMeta() metadomain.Meta {
	currentTimestamp := time.Now().UTC()

	return metadomain.Meta{
		BuildDate:  usecases.metaConfig.BuildDate,
		CommitHash: usecases.metaConfig.CommitHash,
		SystemTime: currentTimestamp,
	}

}
