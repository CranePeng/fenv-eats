package cmd

import (
	"context"
	"fenv-eats/config"
	"fenv-eats/internal/discover"
	"fenv-eats/internal/pool"
	"fenv-eats/internal/service"
	"fenv-eats/models"
	"github.com/CranePeng/fenv-middleware/utils/common"
	l "github.com/CranePeng/fenv-middleware/utils/logger"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"log"
	a "path"
	"runtime"
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
	logger          l.Interface
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
	// 初始化自定义日志(业务日志)
	logger = initMasterLog()
	// 初始化协程池
	pool.Init()
}

func initMasterLog() l.Interface {
	x := common.GetCurrentAbPath()
	x, _ = a.Split(x)
	x = a.Join(x, "logs")
	conf := l.Config{Ctx: context.TODO(), LogLevel: l.Info, CreateFile: true, LogPath: x}
	ll := l.New(conf)
	return ll
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
	err := pool.Work(func() {
		discover.ServiceCluster.WatchNodes(ctx, master.Id)
	})
	if err != nil {
		return
	}
}

func start() {

}
