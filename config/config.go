package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Application struct {
		Company     string
		Project     string
		ServiceName string
		LogLevel    string
		Env         string //环境
		Cluster     string //集群
		Host        string
		Port        int
		Domain      string      //traefik 绑定的域名 （可选）
		AuthUrl     string      //网关统一认证api
		Admins      []UserAdmin //dashboard管理员账号密码
	}
	Etcd struct {
		Endpoints   []string
		DialTimeout int //单位 秒
	}
}

type UserAdmin struct {
	User     string //用户名
	Password string //密码
}

// LoadConfig 加载配置文件
func LoadConfig(file ...string) (*Config, error) {
	fi := "./config.yaml"
	if len(file) > 0 {
		fi = file[0]
	}
	viper.SetConfigFile(fi)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
