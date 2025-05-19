package service

import "github.com/CXeon/micro_contrib/traefik"

var TraefikSvc *traefikService

type traefikService struct {
	*Service
}

func newTraefikSvc(svc *Service) {
	TraefikSvc = &traefikService{svc}
}

// InitDashboard 初始化dashboard
func (svc *traefikService) InitDashboard() error {

	manager, err := traefik.NewManager(svc.etcdCli)
	if err != nil {
		return err
	}
	//组装user
	confAdmins := svc.conf.Application.Admins
	useradmins := make([]traefik.AdminUser, len(confAdmins))

	for i, admin := range confAdmins {
		useradmins[i] = traefik.AdminUser{
			Username: admin.User,
			Password: admin.Password,
		}
	}

	err = manager.InitDashboard(useradmins, svc.conf.Application.Domain)
	if err != nil {
		return err
	}
	return nil
}

// CreateCommonCORS 设置通用跨域中间件
func (svc *traefikService) CreateCommonCORS() error {
	manager, err := traefik.NewManager(svc.etcdCli)
	if err != nil {
		return err
	}

	return manager.CreateCommonCORSMiddleware()
}

// CreateCommonForwardAuthMiddleware 创建通用统一认证中间件
func (svc *traefikService) CreateCommonForwardAuthMiddleware() error {
	manager, err := traefik.NewManager(svc.etcdCli)
	if err != nil {
		return err
	}
	return manager.CreateCommonForwardAuthMiddleware(svc.conf.Application.AuthUrl)
}
