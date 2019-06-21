package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/lexkong/log"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	// 初始化日志包
	c.initLog()
	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}
func (c *Config) initLog() {
    passLagerCfg := log.PassLagerCfg {
		// 输出位置：file/stdout。选择file将日志记录到logger_file 指定的日志文件中，选择 stdout 会将日志输出到标准输出，当然也可以两者同时选择
		Writers:        viper.GetString("log.writers"),
		// 日志级别：DEBUG、INFO、WARN、ERROR、FATAL
		LoggerLevel:    viper.GetString("log.logger_level"),
		// 日志文件
		LoggerFile:     viper.GetString("log.logger_file"),
		// 日志的输出格式，JSON 或者 plaintext，true 会输出成非 JSON 格式，false 会输出成 JSON 格式
		LogFormatText:  viper.GetBool("log.log_format_text"),
		// otate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		// rotate 转存时间，配 合rollingPolicy: daily 使用
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		// rotate 转存大小，配合 rollingPolicy: size 使用
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		// 当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
        LogBackupCount: viper.GetInt("log.log_backup_count"),
    }

    log.InitWithConfig(&passLagerCfg)
}  
// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
