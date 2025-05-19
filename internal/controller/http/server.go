package http

import (
	"context"
	"fmt"
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/traefik_support/config"
	"net/http"
	"time"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 5 * time.Second
)

var httpServer *http.Server

func Start(conf *config.Config, logger *log.Logger) error {
	addr := fmt.Sprintf("%s:%d", conf.Application.Host, conf.Application.Port)
	httpServer = &http.Server{
		Addr:         addr,
		Handler:      initRoutes(conf, logger),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	return httpServer.ListenAndServe()
}

func Stop() error {
	if httpServer == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()
	return httpServer.Shutdown(ctx)
}
