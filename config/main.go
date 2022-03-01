package config

import "os"

type (
	Etcd struct {
		Killer    string   `json:"killer" yaml:"killer" validate:"required"`
		Locker    string   `json:"locker" yaml:"locker" validate:"required"`
		Service   string   `json:"service" yaml:"service" validate:"required"`
		Pipeline  string   `json:"pipeline" yaml:"pipeline" validate:"required"`
		Config    string   `json:"config" yaml:"config" validate:"required"`
		EndPoints []string `json:"endpoints" yaml:"endpoints" validate:"required"`
		Timeout   int64    `json:"timeout" yaml:"timeout" validate:"required"`
	}
	Database struct {
		Host string `json:"host" yaml:"host" validate:"required"`
		Port int    `json:"port" yaml:"port" validate:"required"`
		Name string `json:"name" yaml:"name" validate:"required"`
		User string `json:"user" yaml:"user" validate:"required"`
		Pass string `json:"pass" yaml:"pass" validate:"required"`
		Char string `json:"char" yaml:"char" validate:"required"`
	}
	User struct {
		Name    string `json:"name" yaml:"-" validate:"required"`
		Email   string `json:"email" yaml:"-" validate:"required"`
		Pass    string `json:"pass" yaml:"-" validate:"required"`
		Confirm string `json:"confirm" yaml:"-" validate:"required"`
	}
	Auth struct {
		Secret string `json:"secret" yaml:"secret" validate:"required"`
		TTL    int64  `json:"ttl" yaml:"ttl" validate:"required"`
	}
	Notification struct {
		Url        string `json:"url" yaml:"url" validate:"required"`
		Host       string `json:"host" yaml:"host" validate:"required"`
		Port       int    `json:"port" yaml:"port" validate:"numeric"`
		User       string `json:"user" yaml:"user" validate:"required"`
		Pass       string `json:"pass" yaml:"pass" validate:"required"`
		Name       string `json:"name" yaml:"name" validate:"required"`
		Protocol   string `json:"protocol" yaml:"protocol" validate:"required"`
		Encryption string `json:"encryption" yaml:"encryption" validate:"required"`
	}
	App struct {
		Mode string `json:"mode" yaml:"mode" `
	}
	Docker struct {
		MysqlIp string   `json:"mysqlIp" yaml:"mysqlIp" `
		EtcdIp  []string `json:"etcdIp" yaml:"etcdIp"`
	}
	Config struct {
		Database     `json:"database"`
		Auth         `json:"auth"`
		Etcd         `json:"etcd"`
		Notification `json:"notification"`
		App          `json:"app"`
	}
)

var (
	Conf       *Config
	Path       string
	DockerConf *Docker
)

func Init() *Config {
	DockerConf = &Docker{}
	return &Config{}
}

// 检查配置文件是否存在
func CheckConfigFile(path string) (bool, error) {
	_, err := os.Stat(path)
	exist := !os.IsNotExist(err)
	return exist, err
}
