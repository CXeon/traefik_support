package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/traefik_support/config"
	"github.com/CXeon/traefik_support/internal/controller/http"
	"github.com/CXeon/traefik_support/internal/service"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	httpOri "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Init 初始化并启动服务
func Init() {

	//加载配置文件
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	//创建日志
	logFileName := fmt.Sprintf("./%s.log", conf.Application.ServiceName)
	var logger *log.Logger
	if strings.ToUpper(conf.Application.Env) == "PRO" {
		logger = log.NewLogger(logFileName, zapcore.InfoLevel)
	} else {
		logger = log.NewLogger(logFileName)
	}
	defer logger.Sync()

	logger.Info("server init...")

	//启动etcd客户端

	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Etcd.Endpoints,
		DialTimeout: time.Duration(conf.Etcd.DialTimeout) * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})

	//创建service
	err = service.NewService(conf, logger, etcdCli)
	if err != nil {
		logger.Fatalf("init service err: %s", err.Error())
	}

	//启动web服务
	var eg errgroup.Group
	oriCtx, cancel := context.WithCancel(context.Background())
	eg.Go(func() error {
		logger.Info("start http server")
		err = http.Start(conf, logger)
		if err != nil {
			if !errors.Is(httpOri.ErrServerClosed, err) {
				logger.Errorf("start http err: %s", err.Error())
				cancel()
				return err
			}
			logger.Infof("server exit gracefully")

		}
		return nil
	})

	eg.Go(func() error {
		//优雅退出
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case <-oriCtx.Done():
			return gracefullyExit(etcdCli)
		case <-quit:
			return gracefullyExit(etcdCli)
		}
	})

	err = eg.Wait()
	if err != nil {
		logger.Fatalf("stop server err: %s", err.Error())
	}
	logger.Infof("server stopped")
}

//封装了一些组件的退出流程
func gracefullyExit(etcdCli *clientv3.Client) error {
	var resultErr error
	if etcdCli != nil {
		if err := etcdCli.Close(); err != nil {
			resultErr = err
		}
	}

	if err := http.Stop(); err != nil {
		resultErr = err
	}

	return resultErr
}
