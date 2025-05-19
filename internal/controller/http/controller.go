package http

import (
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/traefik_support/config"
)

type controller struct {
	conf   *config.Config
	logger *log.Logger
}

func NewController(conf *config.Config, logger *log.Logger) *controller {
	if conf == nil || logger == nil {
		return nil
	}

	return &controller{
		conf:   conf,
		logger: logger,
	}
}
