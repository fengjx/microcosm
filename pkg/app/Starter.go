package app

import (
	"microcosm/conf"
)

type Starter interface {
	Start(config *conf.Config)
	Stop(config *conf.Config)
}
