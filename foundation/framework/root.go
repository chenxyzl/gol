package framework

import (
	"encoding/base64"
	"fmt"
	"foundation/framework/config"
	"foundation/framework/g"
	"foundation/framework/log"
	"foundation/framework/version"
	"foundation/table"
	"github.com/spf13/cobra"
	"strings"
)

var (
	cfgFile  string
	tableDir string
	nodeId   uint64
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:          "foundation",
	Short:        "foundation server",
	SilenceUsage: true,
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
		fmt.Println("BuildTime:", version.BuildTime)
		fmt.Println("CommitID:", version.CommitID)
		fmt.Println("GoVer:", version.GoVer)
		x := strings.ReplaceAll(version.VersionMsg, "#", "=")
		if bts, err := base64.StdEncoding.DecodeString(x); err == nil {
			fmt.Print(string(bts))
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "serverConfig", "conf/dev_cfg.toml", "server atom config path")
	RootCmd.PersistentFlags().StringVar(&tableDir, "tableDir", "", "table path")
	RootCmd.PersistentFlags().Uint64Var(&nodeId, "id", 0, "node id(must unique in same cluster)")
	cobra.OnInitialize(onInit)
}

func onInit() {
	//先初始化游戏配置
	config.Init(cfgFile)
	//全局初始化
	g.Init(nodeId)
	//在初始化日志
	log.Init()
	//初始化配置表
	table.Init()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute(cmds ...*cobra.Command) {
	//添加命令
	for _, cmd := range cmds {
		RootCmd.AddCommand(cmd)
	}
	RootCmd.AddCommand(versionCmd)
	//开始执行
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
