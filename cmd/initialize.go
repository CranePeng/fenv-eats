package cmd

import (
	"context"
	"encoding/json"
	"fenv-eats/config"
	"fenv-eats/internal/discover"
	"fenv-eats/internal/service"
	"fenv-eats/internal/utils"
	"fenv-eats/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	clientV3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// installCmd represents the installed command
// 初始化命令，只是用来初始化项目配置，启动项目请看master和worker的命令
var (
	initializeCmd = &cobra.Command{
		Use:     "init",
		Short:   "Run initialize elastic automatic task system",
		Long:    "Run initialize elastic automatic task system",
		Example: "eats init",
		Run: func(cmd *cobra.Command, args []string) {
			switch mode {
			case "web":
				startInitializeWeb()
				break
			case "json":
				startInitializeByConfigFile()
			case "yaml":
				startInitializeByConfigFile()
				break
			default:
				log.Println("Unsupported file extension")
				os.Exit(1)
			}
		},
	}

	mode string
	path string

	user = &models.User{
		Id:         uuid.NewV4().String(),
		Manager:    true,
		CreateTime: utils.Time(time.Now()),
		UpdateTime: utils.Time(time.Now()),
	}
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.Conf = config.Init()
	rootCmd.AddCommand(initializeCmd)
	service.Runtime = &service.Instance{
		Version: rootCmd.Version,
	}
	initializeCmd.Flags().StringVarP(&mode, "mode", "m", "web", "Set initialize mode with web ui or json, yaml config file")
	initializeCmd.Flags().StringVarP(&path, "path", "p", "", "Set config file path")
	initializeCmd.Flags().StringVarP(&user.Name, "name", "n", "", "Set admin name")
	initializeCmd.Flags().StringVarP(&user.Email, "email", "e", "", "Set admin email")
	initializeCmd.Flags().StringVarP(&user.Password, "pass", "P", "", "Set admin pass")

}

func startInitializeWeb() {
	app := gin.New()
	app.Use(gin.Recovery())

	sg := new(sync.WaitGroup)
	defer sg.Wait()

}

func startInitializeByConfigFile() {
	validatePath()
	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Fatal("Please enter admin user info")
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if mode == "json" {
		if err := json.Unmarshal(buf, config.Conf); err != nil {
			log.Fatal(err)
		}
	}

	if mode == "yaml" {
		if err := yaml.Unmarshal(buf, config.Conf); err != nil {
			log.Fatal(err)
		}
	}
	// 初始化etcd客户端
	discover.NewClient()

	buf, err = json.Marshal(config.Conf)
	if err != nil {
		log.Fatal(err)
	}

	if res, err := discover.Client.Put(context.TODO(), config.Conf.Etcd.Config, string(buf), clientV3.WithPrevKV()); err != nil {
		log.Fatal(err)
	} else {
		if len(res.PrevKv.Value) > 0 {
			log.Printf("%s %s \n", "OLD CONFIG IS", string(res.PrevKv.Value))
		}
	}
	// 初始化数据库：配置初始用户
	initDatabase()
	log.Println("INITIALIZE SUCCESSFUL")
}

func initDatabase() {
	var err error
	if err = utils.CreateDatabase(); err != nil {
		log.Fatal("Failed to create database", err)
	}

	if models.Engine == nil {
		// Create database engine
		err = models.Connection()
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
	}

	if err := models.Migrate(); err != nil {
		log.Fatal("Failed to migrate the table", err)
	}

	pass, err := models.GeneratePassword(user.Password)
	user.Password = string(pass)

	if _, err := models.Engine.Insert(user); err != nil {
		log.Fatal("Failed to create system manager", err)
	}
}

func validatePath() {
	if path == "" {
		log.Fatal("Please enter your config file path or use --mode=web")
	}

	exist, err := config.CheckConfigFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if exist == false {
		log.Fatal("Config file does not exist")
	}
}
