package server

import (
	"fmt"
	"github.com/cuiyuanxin/airuisi-admin/global"
	"github.com/cuiyuanxin/airuisi-admin/pkg/server"
	"github.com/spf13/cobra"
	"github.com/zztroot/color"
	"log"
	"os"
)

var (
	envConf string
	Cmd     = &cobra.Command{
		Use:     "server",
		Short:   "Starting the server",
		Long:    "The default config configuration can be overridden by specifying the config file path, which is optional",
		Example: global.AppName + " server -c config/config.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	envConf = os.Getenv("APP_CONF")
	if envConf == "" {
		Cmd.PersistentFlags().StringVarP(&envConf, "config", "c", "config/prod.yaml", "config path, eg: -conf config/prod.yml")
	}
	if envConf == "" {
		envConf = "config/prod.yaml"
	}
	log.Println("Load the conf file:", envConf)
}

func setup() {
	// 初始化配置
	if err := setupSetting(envConf); err != nil {
		log.Fatalf("setup.setupSetting err：%s", color.Coat(err.Error(), color.Red))
	}
	log.Println(color.Coat("Configuration initialization successfully...", color.Green))
	if err := setupLogger(); err != nil {
		log.Fatalf("setup.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	}
	log.Println(color.Coat("Log initialization was successful...", color.Green))
	//err = setupDBEngine()
	//if err != nil {
	//	log.Fatalf("init.setupDBEngine err：%s", color.Coat(err.Error(), color.Red))
	//}

	//err = setupTracer()
	//if err != nil {
	//	log.Fatalf("init.setupTracet err：%s", color.Coat(err.Error(), color.Red))
	//}
	//err = setupTranslations()
	//if err != nil {
	//	log.Fatalf("init.setupTranslations err：%s", color.Coat(err.Error(), color.Red))
	//}
	//err = setupCasbin()
	//if err != nil {
	//	log.Fatalf("init.setupCasbin err：%s", color.Coat(err.Error(), color.Red))
	//}

	// 初始化日志
	//if err := logger.NewLogger(global.LoggerSetting); err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//log.Println("Log initialization successfully...")

	// 初始化链路追踪
	//if err := tracer.NewJaegerTracer(global.AppName+` server`, global.TracerSetting.AgentHostPort); err != nil {
	//	log.Fatalf("init.setupTracet err：%s", color.Coat(err.Error(), color.Red))
	//}
	//log.Println("Link tracking was initialized successfully...")

	// 初始化数据库
	//if err := model.NewDBEngine(global.DatabaseSetting); err != nil {
	//	log.Fatalf("init.setupDBEngine err：%s", color.Coat(err.Error(), color.Red))
	//}
	//log.Println("Database initialization complete...")

}

func tip() {
	usageStr := `欢迎使用 ` + color.Coat(fmt.Sprintf("%s %s", global.AppName, global.Version), color.Green) + ` 可以使用 ` + color.Coat("-h", color.Red) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func run() error {

	app, err := newApp(global.Logger)
	if err != nil {
		log.Panicf("run.server err：%s", color.Coat(err.Error(), color.Red))
	}

	tip()
	fmt.Println(color.Coat("Server run at:", color.Green))
	fmt.Printf("-  Local:   %s://localhost:%s/ \r\n", "http", global.ServerSetting.Port)
	fmt.Printf("-  Network: %s://%s:%s/ \r\n", "http", global.ServerSetting.Addr, global.ServerSetting.Port)
	fmt.Println(color.Coat("Swagger run at:", color.Green))
	//fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", config.ApplicationConfig.Port)
	//fmt.Printf("-  Network: %s://%s:%d/swagger/admin/index.html \r\n", "http", pkg.GetLocaHonst(), config.ApplicationConfig.Port)
	//fmt.Printf("%s Enter Control + C Shutdown Server \r\n", pkg.GetCurrentTimeStr())

	server.Run(app)

	return nil
}
