package models

import (
	"fenv-eats/config"
	"fenv-eats/log"
	"fmt"
	orm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	Engine    *orm.DB
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
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s",
			config.Conf.Database.User,
			config.Conf.Database.Pass,
			config.Conf.Database.Host,
			config.Conf.Database.Port,
			config.Conf.Database.Name,
			config.Conf.Database.Char,
		)
		Engine, err = orm.Open("mysql", dsn)
		if err != nil {
			errR = err
		}
		// 连接数的配置也可以在这里写，也可以写入配置文件
		Engine.DB().SetMaxIdleConns(50)
		Engine.DB().SetMaxOpenConns(50)
		Engine.SetLogger(log.Instance)
		Engine.LogMode(true)
		// 自动迁移数据结构(table schema)
		//注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
		//db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
		Engine.SingularTable(true)
	})
	if errR != nil {
		return errR
	}
	return nil

}
