package config

import (
	"fmt"
	"foundation/framework/g"
	"foundation/framework/version"
	"foundation/framework/version_impl"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gitlab-ee.funplus.io/watcher/watcher/wrpc/client"
	"os"
	"path/filepath"
	"strings"
)

// 获取工程的基础目录.譬如k-server绝对地址
func getBaseDir() (string, error) {
	exePath := os.Args[0]
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(absPath)
	if strings.HasSuffix(dir, fmt.Sprintf("%cbin", os.PathSeparator)) {
		dir = filepath.Dir(dir)
	} else if strings.HasSuffix(filepath.Dir(dir), fmt.Sprintf("%cbin", os.PathSeparator)) {
		dir = filepath.Dir(dir)
	}
	return dir, nil
}

func Init(cfgFile string) {
	g.GlobalConfig = viper.New()
	config := g.GlobalConfig
	if cfgFile != "" {
		// 判断绝对路径
		if strings.HasPrefix(cfgFile, "/") {
			config.SetConfigFile(cfgFile)
		} else {
			BaseDir, err := getBaseDir()
			fmt.Printf("Base dir<%s>\n", BaseDir)
			if err != nil {
				panic(err)
			}
			fmt.Println("cfgFile is :", cfgFile)
			// Use config file from the flag.
			cfgFile = BaseDir + "/" + cfgFile
			config.SetConfigFile(cfgFile)
			fmt.Println("Full cfg file path:", cfgFile)
		}

	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".json" (without extension).
		config.AddConfigPath(home)
		config.SetConfigName("foundation_cfg")
		config.SetConfigType("toml")
	}

	config.AutomaticEnv() // read in environment variables that match
	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", config.ConfigFileUsed())
		panic(err)
	}

	g.PProf = config.GetString("PPROF.Addr")
	g.EtcdAddr = config.GetString("ETCD.Addr")
	g.DebugModel = config.GetBool("Global.DebugMode")
	g.EnableWdsyncMD5Map = config.GetBool("Kingdom.enableWdsyncMD5Map")
	version.VersionMng, _ = version_impl.NewVersionMng(config.GetString("Version.kingdom"))
	version.VersionMng.SetClientVersion(config.GetString("Version.client"))
	client.DefaultConnOptsB.ReadTimeout = config.GetDuration("Global.ConnReadTimeoutB")
}
