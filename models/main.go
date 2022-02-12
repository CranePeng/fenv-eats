package models

import (
	"fenv-eats/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	Engine    *gorm.DB
	err       error
	mysqlOnce sync.Once
)

type (
	// 用于填充系统默认数据的接口
	Seeder interface {
		Seed() error
	}
	// 模型接口
	Model interface {
		Store() error
		Update() error
		ToString() (string, error)
	}
)

const DefaultTimeFormat = "2006-01-02 15:04:05"

func Connection() error {
	var errR error
	mysqlOnce.Do(func() {
		mysqlIp := config.Conf.Database.Host
		if len(config.DockerConf.MysqlIp) > 0 {
			mysqlIp = config.DockerConf.MysqlIp
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s",
			config.Conf.Database.User,
			config.Conf.Database.Pass,
			mysqlIp,
			config.Conf.Database.Port,
			config.Conf.Database.Name,
			config.Conf.Database.Char,
		)
		Engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			errR = err
		}
		sqlDB, err := Engine.DB()
		if err != nil {
			errR = err
		}
		// 连接数的配置也可以在这里写，也可以写入配置文件
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(50)
		// SetMaxOpenConns 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(50)
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(time.Hour)

		// 自动迁移数据结构(table schema)
		//注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
		//db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致,2.0已经移除
		//Engine.SingularTable(true)
	})
	if errR != nil {
		return errR
	}
	return nil

}

// 迁移数据库
func Migrate() error {
	tables := []interface{}{
		&User{},
		&Node{},
		&Task{},
		&Log{},
		&PasswordResets{},
		&Pipeline{},
		&PipelineRecords{},
		&PipelineTaskPivot{},
		&PipelineNodePivot{},
		&TaskRecords{},
	}

	Engine.AutoMigrate(tables...)
	return nil
}
