package cmd

import (
	"context"
	"fenv-eats/config"
	"fenv-eats/internal/discover"
	"fenv-eats/internal/pool"
	"fenv-eats/internal/service"
	"fenv-eats/models"
	"fenv-eats/routes"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

// masterCmd represents the master command
var (
	masterCmd = &cobra.Command{
		Use:   "master",
		Short: "Run a master node service",
		Long:  "Run a master node service on this server",
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap()
			watch()
			start()
		},
	}

	master = &service.Instance{
		Mode:    models.MASTER,
		Status:  models.ONLINE,
		Version: rootCmd.Version,
	}

	ctx, cancelFunc = context.WithCancel(context.Background())
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.AddCommand(masterCmd)
	config.Conf = config.Init()
	service.Initialize()
	masterCmd.Flags().StringVar(&master.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.Flags().IntVar(&master.Port, "port", 9402, "Set listen on port")
	masterCmd.Flags().StringSliceVar(&service.EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	masterCmd.Flags().StringVarP(&master.Id, "node", "n", uuid.NewV4().String(), "Set master node id")
	masterCmd.Flags().StringVar(&master.Name, "name", "", "Set master node name")
	masterCmd.Flags().StringVar(&master.Description, "desc", "master node", "Set master node description")
	masterCmd.Flags().StringVar(&service.ConfigKey, "config", "/eagle/config", "Set the key used to get configuration information")
	masterCmd.Flags().StringVar(&config.DockerConf.MysqlIp, "mysql", "127.0.0.1", "master mysql ip")
}

func bootstrap() {
	var err error
	config.Conf.Etcd.EndPoints = service.EndPoints
	discover.NewClient()
	discover.GetConf(service.ConfigKey)
	err = models.Connection()
	if err != nil {
		log.Fatal(err)
	}
}

func watch() {
	err := pool.GetWorkerInstance().Work(func() {
		cluster := &discover.Cluster{
			Client: discover.Client,
		}
		cluster.WatchNodes(ctx, master.Id)
	})
	if err != nil {
		return
	}
}

func start() {
	g := gin.New()
	g = routes.Init(g)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(master.Port),
		Handler: g,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在此阻塞
	<-quit

	ctx, channelCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channelCancel()

	// 当主节点关闭时，先检查是否有其他在线主节点

	// 关闭主节点的上下文
	cancelFunc()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}
	log.Println("server exiting...")

}
