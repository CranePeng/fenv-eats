package cmd

import (
	"context"
	"github.com/CranePeng/fenv-middleware/utils/common"
	l "github.com/CranePeng/fenv-middleware/utils/logger"
	"github.com/spf13/cobra"
	"os"
	a "path"
)

var rootCmd = &cobra.Command{
	Use:     "Eagle",
	Short:   "Eagle",
	Long:    "Eagle:Elastic Automatic Task System",
	Version: "0.1",
}
var Logger l.Interface

func Execute() {
	// 初始化自定义日志(业务日志)
	Logger = initMasterLog()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initMasterLog() l.Interface {
	x := common.GetCurrentAbPath()
	x, _ = a.Split(x)
	x = a.Join(x, "logs")
	conf := l.Config{Ctx: context.Background(), LogLevel: l.Info, CreateFile: true, LogPath: x}
	ll := l.New(conf)
	return ll
}
