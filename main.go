package main

import (
	"clients/config"
	"clients/middleware"
	"clients/pkg/cron"
	"clients/pkg/token"
	v "clients/pkg/version"
	"clients/pkg/zerodown"
	"clients/routers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path")
	version = pflag.BoolP("version", "v", false, "show version info.")
	tk      = pflag.BoolP("token", "t", false, "shaking hands with the master.")
)

// @title Web API
// @version 1.0
// @description Web ops

// @contact.name guang
// @contact.url http://www.swagger.io/support
// @contact.email av1254@qq.com

// @host localhost:8080
// @BasePath /v1
func main() {

	pflag.Parse()
	if *version {
		v := v.Get()
		marshaled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshaled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	if *tk {
		err := token.Create()
		if err != nil {
			fmt.Println(`添加主机失败`, err)
		}
		return
	}

	// init cron
	go func() {
		cron.Init()
		defer cron.Close()
	}()

	os.MkdirAll(viper.GetString("shell.log"), os.ModePerm)

	gin.SetMode(viper.GetString("runmode"))

	//  Create gin
	g := gin.New()

	routers.Load(
		g,
		//middleware...,
		middleware.Logging(),
		middleware.RequestID(),
	)
	//go func() {
	//	if err := pingServer(); err != nil {
	//		log.Fatal("The router has no respnese, or it might took too long to start up", err)
	//	}
	//	log.Info("The router has been deployed successfully.")
	//}()

	// Start to listening the incoming requests
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("start to listening the incoming requests on http address: %s", viper.GetString("addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	if viper.GetString("runmode") == "debug" {
		log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
	} else {
		log.Info(zerodown.ListenAndServe(viper.GetString("addr"), g).Error()) // 编译时打开 注意: win下不可用
	}
	log.Infof("start to listening the incoming requests on http address: %s", viper.GetString("addr"))
}

//func pingServer() error {
//	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
//		resp, err := http.Get(viper.GetString("url") + "/test/health")
//		if err == nil && resp.StatusCode == 200 {
//			return nil
//		}
//
//		log.Info("Waiting for the router, retry in 1 second.")
//		time.Sleep(time.Second)
//	}
//	return errors.New("Cannot connect to the router")
//}
