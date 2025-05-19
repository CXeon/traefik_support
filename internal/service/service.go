package service

import (
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/traefik_support/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Service struct {
	conf    *config.Config
	logger  *log.Logger
	etcdCli *clientv3.Client
}

func NewService(conf *config.Config, logger *log.Logger, etcdClient *clientv3.Client) error {
	svc := &Service{
		conf:    conf,
		logger:  logger,
		etcdCli: etcdClient,
	}

	newTraefikSvc(svc)

	//初始化dashboard和中间件
	var err error
	err = TraefikSvc.InitDashboard()
	err = TraefikSvc.CreateCommonCORS()
	err = TraefikSvc.CreateCommonForwardAuthMiddleware()

	return err
}
