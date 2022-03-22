package log

import (
	"foundation/framework/g"
	"gitlab-ee.funplus.io/watcher/watcher/misc"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"gitlab-ee.funplus.io/watcher/watcher/misc/zap"
	"gitlab-ee.funplus.io/watcher/watcher/misc/zap/zapcore"
)

type LogConfig struct {
	logFile              bool
	logLevel             string
	conditionLogLevel    string
	logProd              bool
	logTag               string
	maxLogNum            int
	maxSizeofEachLogFile int64
	isShowInConsole      bool
}

func Init() {
	config := g.GlobalConfig
	logConfig := &LogConfig{}
	logConfig.logFile = config.GetBool("Log.LogFile")
	logConfig.logProd = config.GetBool("Log.LogProd")
	logConfig.logTag = config.GetString("Log.LogTag")
	logConfig.logLevel = config.GetString("Log.LogLevel")
	logConfig.conditionLogLevel = config.GetString("Log.ConditionLogLevel")
	logConfig.maxLogNum = config.GetInt("Log.maxlognum")
	logConfig.maxSizeofEachLogFile = config.GetInt64("Log.maxsizeofeachlogfile")
	config.SetDefault("Log.isshowinconsole", true)
	logConfig.isShowInConsole = config.GetBool("Log.isshowinconsole")

	ip, err := misc.GetIP("10.*", "192.*", "172.*")
	if err != nil {
		panic(err)
	}

	// Init the log module
	var fields = []zapcore.Field{
		//zap.String("me", name),
	}
	if logConfig.logProd {
		fields = append(fields, zap.String("serverIP", ip.String()))
	}
	if len(logConfig.logTag) > 0 {
		fields = append(fields, zap.String("tag", logConfig.logTag))
	}

	if logConfig.logFile {
		err := wlog.InitFileLog("log/debuglog.txt",
			wlog.WithOptionServerName(string(g.RoleTyp())),
			wlog.WithOptionLogLevel(logConfig.logLevel),
			wlog.WithOptionConditionLogLevel(logConfig.conditionLogLevel),
			wlog.WithOptionMaxLogNum(logConfig.maxLogNum),
			wlog.WithOptionMaxLogSize(logConfig.maxSizeofEachLogFile),
			wlog.WithOptionShowLogInConsole(logConfig.isShowInConsole))
		if err != nil {
			panic(err)
		}
		var logCfg = wlog.NewDevelopmentFileLogConfig(wlog.ConfigWithIgnoreLogPath("foundation/logs/"))
		if err := wlog.InitializeFileLog(logCfg, zap.Fields(fields...)); err != nil {
			panic(err)
		}
	} else {
		err := wlog.InitNoFileLog(wlog.WithOptionServerName(string(g.RoleTyp())),
			wlog.WithOptionLogLevel(logConfig.logLevel),
			wlog.WithOptionConditionLogLevel(logConfig.conditionLogLevel),
			wlog.WithOptionShowLogInConsole(logConfig.isShowInConsole))
		if err != nil {
			panic(err)
		}
		var logCfg = wlog.NewDevelopmentConfig(wlog.ConfigWithIgnoreLogPath("foundation/logs/"))
		if err := wlog.InitializeNoFileLog(logCfg, zap.Fields(fields...)); err != nil {
			panic(err)
		}
	}

	wlog.SetDebugMode(g.DebugModel)
}
