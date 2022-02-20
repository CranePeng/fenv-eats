# fenv-eagle

> eagle，全称：elastic automatic task system (eats) 是分布式任务调度中心，前身参考 [ects](https://betterde.github.io/ects/) ,在此框架设计基础上补全了追踪和资源回收优化
> ，不同于ects的是，eagle采用了gin+gorm+redis，补全了协程池和任务监控，于是改名成了eagle


```go
// 依赖包

// gin框架相关
go get github.com/dchest/captcha
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/redis

// 日志框架
go get github.com/sirupsen/logrus
go get github.com/natefinch/lumberjack
go get github.com/lestrrat-go/file-rotatelogs
go get github.com/antonfisher/nested-logrus-formatter

// orm 框架
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
go get github.com/go-redis/redis/v8

// 性能监控
go get github.com/gin-contrib/pprof

// 命令行
go get github.com/spf13/cobra
// uuid
go get github.com/satori/go.uuid

// 表单校验
go get github.com/go-playground/validator/v10

// etcd 作为注册中心和服务发现中间件
go get go.etcd.io/etcd/client/v3
go get github.com/coreos/etcd/clientv3

// 协程池
go get -u github.com/panjf2000/ants/v2

// 自研中间件（链路追踪相关）
go get github.com/CranePeng/fenv-middleware

```

# 启动 MySQL
docker run -d \
--name mysql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=your-secret-pw mysql:8.0.21 \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_unicode_ci

# 启动 ETCD
$ docker run -d \
--name etcd \
-p 2379:2379 \
-p 2380:2380 \
--volume=/Users/pengzihe/Developer/GoProjects/fenv/fenv-eats/data/etcd:/data \
gcr.io/etcd-development/etcd:v3.5.1 \
/usr/local/bin/etcd \
--name etcd \
--data-dir /data \
--listen-client-urls http://0.0.0.0:2379 \
--advertise-client-urls http://0.0.0.0:2379 \
--listen-peer-urls http://0.0.0.0:2380 \
--initial-advertise-peer-urls http://0.0.0.0:2380 \
--initial-cluster etcd=http://0.0.0.0:2380 \
--initial-cluster-token fenv \
--initial-cluster-state new

docker run --network=fenv -d --name fenv-demo2 --link etcd:etcd -p 9701:9701 ects:1.0



## 项目编译

go build -ldflags="-s -w -H windowsgui" maim.go -o main.exe

    1.-s strip 去掉无用的符号
    2.-w DWARF 去掉DWARF调试信息，得到的可执行程序不可用调试器调试
    3.-H windowsgui 生成带GUI界面的程序时，可去掉dos黑框
    
    以上为3个常用的参数，此外-ldflags '-extldflags "-static"' 为静态编译
    如果，想更加清楚的看到编译过程可加-x 参数，如bulid -x ......
    
    
    
    -H windowsgui隐藏go自己输出的命令行窗口， 隐藏调用的外部程序的cmd窗口

最终：
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "ects_linux" main.go





