package app

import (
	"microcosm/internal/config"
)

type Starter interface {
	Start(config *config.Config)
	Stop(config *config.Config)
}
