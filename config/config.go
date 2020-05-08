package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// init
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件的变化，并热加载
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 解析指定文件
		viper.SetConfigFile(c.Name)
	} else {
		// 没有指定文件，就去拿默认文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件的格式为yaml
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_leverl"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed %s", e.Name)
	})
}
